package models_test

import (
	"fmt"
	"go/parser"
	"go/token"
	"testing"

	"github.com/fullstack-lang/gong/dsm/process/go/models"
)

func TestGenerateOrderToCashModel(t *testing.T) {
	stage := models.NewStage("test")

	// 1. Root Library
	rootLibrary := (&models.Library{
		Name:              "Order to Cash",
		IsRootLibrary:     true,
		NbPixPerCharacter: 8.5,
	}).Stage(stage)

	// 2. Resources
	crmResource := (&models.Resource{Name: "CRM System"}).Stage(stage)
	erpResource := (&models.Resource{Name: "ERP System"}).Stage(stage)
	wmsResource := (&models.Resource{Name: "WMS (Warehouse Management)"}).Stage(stage)
	paymentResource := (&models.Resource{Name: "Payment Gateway"}).Stage(stage)
	rootLibrary.RootResources = append(rootLibrary.RootResources, crmResource, erpResource, wmsResource, paymentResource)

	// 3. Data items
	poData := (&models.Data{Name: "Purchase Order", Description: "Customer Purchase Order with line items and pricing"}).Stage(stage)
	ackData := (&models.Data{Name: "Order Confirmation", Description: "Sales order confirmation with promised shipping schedule"}).Stage(stage)
	pickData := (&models.Data{Name: "Picking List", Description: "Itemized warehouse pick and pack instruction slip"}).Stage(stage)
	bolData := (&models.Data{Name: "Bill of Lading", Description: "Freight waybill and carrier consignment tracking record"}).Stage(stage)
	invData := (&models.Data{Name: "Commercial Invoice", Description: "Tax and commercial invoice specifying payment terms"}).Stage(stage)
	remitData := (&models.Data{Name: "Payment Remittance", Description: "Electronic funds transfer confirmation and remittance advice"}).Stage(stage)
	rootLibrary.RootDatas = append(rootLibrary.RootDatas, poData, ackData, pickData, bolData, invData, remitData)

	// 4. Process
	o2cProcess := (&models.Process{
		Name:        "Order to Cash",
		Description: "Standard Order-to-Cash (O2C) business process orchestrating customer order placement, order validation, inventory fulfillment, shipping, invoicing, and payment reconciliation across multiple enterprise actors.",
	}).Stage(stage)
	rootLibrary.RootProcesses = append(rootLibrary.RootProcesses, o2cProcess)

	// 5. External Participants
	customer := (&models.Participant{
		Name:        "Customer",
		Description: "External buying customer placing orders and settling commercial invoices.",
	}).Stage(stage)

	carrier := (&models.Participant{
		Name:        "Logistics Carrier",
		Description: "Third-party freight carrier handling parcel dispatch and physical delivery.",
	}).Stage(stage)

	o2cProcess.ExternalParticipants = append(o2cProcess.ExternalParticipants, customer, carrier)

	// 6. Internal Participants (Swimlanes)
	sales := (&models.Participant{
		Name:        "Sales & Order Desk",
		Description: "Internal sales operations validating customer orders, checking credit, and confirming sales agreements.",
		Resources:   []*models.Resource{crmResource},
	}).Stage(stage)

	warehouse := (&models.Participant{
		Name:        "Warehouse & Logistics",
		Description: "Fulfillment center checking inventory, picking, packing, and dispatching products to carriers.",
		Resources:   []*models.Resource{wmsResource, erpResource},
	}).Stage(stage)

	finance := (&models.Participant{
		Name:        "Finance & Billing",
		Description: "Accounting department generating legal commercial invoices and reconciling customer payment remittances.",
		Resources:   []*models.Resource{erpResource, paymentResource},
	}).Stage(stage)

	o2cProcess.Participants = append(o2cProcess.Participants, sales, warehouse, finance)

	// 7. Tasks
	// Sales tasks
	startSales := (&models.Task{Name: "Start Order Desk", IsStartTask: true}).Stage(stage)
	validateOrder := (&models.Task{Name: "Receive & Validate Order", Description: "Check customer credentials, item prices, and credit limits"}).Stage(stage)
	confirmOrder := (&models.Task{Name: "Confirm Sales Order", Description: "Issue formal sales order acknowledgment to customer and release to logistics"}).Stage(stage)
	endSales := (&models.Task{Name: "End Order Desk", IsEndTask: true}).Stage(stage)
	sales.Tasks = append(sales.Tasks, startSales, validateOrder, confirmOrder, endSales)

	// Warehouse tasks
	startWh := (&models.Task{Name: "Start Logistics", IsStartTask: true}).Stage(stage)
	reserveStock := (&models.Task{Name: "Check & Reserve Inventory", Description: "Verify on-hand warehouse stock and allocate inventory"}).Stage(stage)
	pickPack := (&models.Task{Name: "Pick & Pack Products", Description: "Retrieve items from bins and package into shipping parcels"}).Stage(stage)
	dispatchGoods := (&models.Task{Name: "Dispatch & Ship Goods", Description: "Affix carrier waybill labels and transfer parcels to carrier"}).Stage(stage)
	endWh := (&models.Task{Name: "End Logistics", IsEndTask: true}).Stage(stage)
	warehouse.Tasks = append(warehouse.Tasks, startWh, reserveStock, pickPack, dispatchGoods, endWh)

	// Finance tasks
	startFin := (&models.Task{Name: "Start Finance", IsStartTask: true}).Stage(stage)
	issueInvoice := (&models.Task{Name: "Generate Commercial Invoice", Description: "Issue electronic tax invoice based on shipped quantities"}).Stage(stage)
	processPayment := (&models.Task{Name: "Process Payment & Settle Account", Description: "Match incoming remittance with invoice and clear accounts receivable"}).Stage(stage)
	closeOrder := (&models.Task{Name: "Close Order & Archive Record", IsEndTask: true, Description: "Mark order as fully fulfilled and archive audit trail"}).Stage(stage)
	finance.Tasks = append(finance.Tasks, startFin, issueInvoice, processPayment, closeOrder)

	// 8. Control Flows
	createCF := func(p *models.Participant, start, end *models.Task) *models.ControlFlow {
		cf := (&models.ControlFlow{
			Name:  fmt.Sprintf(`"%s" to "%s"`, start.Name, end.Name),
			Start: start,
			End:   end,
		}).Stage(stage)
		p.ControlFlows = append(p.ControlFlows, cf)
		return cf
	}

	cfSales1 := createCF(sales, startSales, validateOrder)
	cfSales2 := createCF(sales, validateOrder, confirmOrder)
	cfSales3 := createCF(sales, confirmOrder, endSales)

	cfWh1 := createCF(warehouse, startWh, reserveStock)
	cfWh2 := createCF(warehouse, reserveStock, pickPack)
	cfWh3 := createCF(warehouse, pickPack, dispatchGoods)
	cfWh4 := createCF(warehouse, dispatchGoods, endWh)

	cfFin1 := createCF(finance, startFin, issueInvoice)
	cfFin2 := createCF(finance, issueInvoice, processPayment)
	cfFin3 := createCF(finance, processPayment, closeOrder)

	// 9. Data Flows
	// Ext to Task: Customer -> validateOrder (PO)
	df1 := (&models.DataFlow{
		Name:                     "Customer to Receive & Validate Order",
		Type:                     models.DataFlow_ExternalParticipant2Task,
		StartExternalParticipant: customer,
		EndTask:                  validateOrder,
		Datas:                    []*models.Data{poData},
	}).Stage(stage)

	// Task to Ext: confirmOrder -> Customer (Ack)
	df2 := (&models.DataFlow{
		Name:                   "Confirm Sales Order to Customer",
		Type:                   models.DataFlow_Task2ExternalParticipant,
		StartTask:              confirmOrder,
		EndExternalParticipant: customer,
		Datas:                  []*models.Data{ackData},
	}).Stage(stage)

	// Task to Task: confirmOrder -> reserveStock (Ack)
	df3 := (&models.DataFlow{
		Name:      `"Confirm Sales Order" to "Check & Reserve Inventory"`,
		Type:      models.DataFlow_Task2Task,
		StartTask: confirmOrder,
		EndTask:   reserveStock,
		Datas:     []*models.Data{ackData},
	}).Stage(stage)

	// Task to Ext: dispatchGoods -> Carrier (BOL)
	df4 := (&models.DataFlow{
		Name:                   "Dispatch & Ship Goods to Logistics Carrier",
		Type:                   models.DataFlow_Task2ExternalParticipant,
		StartTask:              dispatchGoods,
		EndExternalParticipant: carrier,
		Datas:                  []*models.Data{bolData},
	}).Stage(stage)

	// Task to Task: dispatchGoods -> issueInvoice (BOL)
	df5 := (&models.DataFlow{
		Name:      `"Dispatch & Ship Goods" to "Generate Commercial Invoice"`,
		Type:      models.DataFlow_Task2Task,
		StartTask: dispatchGoods,
		EndTask:   issueInvoice,
		Datas:     []*models.Data{bolData},
	}).Stage(stage)

	// Task to Ext: issueInvoice -> Customer (Invoice)
	df6 := (&models.DataFlow{
		Name:                   "Generate Commercial Invoice to Customer",
		Type:                   models.DataFlow_Task2ExternalParticipant,
		StartTask:              issueInvoice,
		EndExternalParticipant: customer,
		Datas:                  []*models.Data{invData},
	}).Stage(stage)

	// Ext to Task: Customer -> processPayment (Remittance)
	df7 := (&models.DataFlow{
		Name:                     "Customer to Process Payment & Settle Account",
		Type:                     models.DataFlow_ExternalParticipant2Task,
		StartExternalParticipant: customer,
		EndTask:                  processPayment,
		Datas:                    []*models.Data{remitData},
	}).Stage(stage)

	allDataFlows := []*models.DataFlow{df1, df2, df3, df4, df5, df6, df7}
	rootLibrary.RootDataFlows = append(rootLibrary.RootDataFlows, allDataFlows...)
	o2cProcess.DataFlows = append(o2cProcess.DataFlows, allDataFlows...)

	// 10. Diagrams
	// Helper function for creating shapes in Diagram 1
	createDiagram1 := func() *models.DiagramProcess {
		diag := (&models.DiagramProcess{
			Name:             "Order to Cash - Full Process",
			Description:      "End-to-end view of the Order-to-Cash process across Sales, Logistics, and Finance with Customer and Carrier touchpoints.",
			IsChecked:        true,
			IsEditable_:      true,
			DefaultBoxWidth:  240,
			DefaultBoxHeigth: 65,
			Width:            1650,
			Height:           1080,
		}).Stage(stage)

		// ProcessShape
		pShape := (&models.ProcessShape{
			Name:    "ProcessShape",
			Process: o2cProcess,
		}).Stage(stage)
		pShape.SetX(250)
		pShape.SetY(40)
		pShape.SetWidth(1060)
		pShape.SetHeight(990)
		diag.Process_Shapes = append(diag.Process_Shapes, pShape)

		// Participant Shapes
		makePShape := func(p *models.Participant) *models.ParticipantShape {
			ps := (&models.ParticipantShape{
				Name:        fmt.Sprintf("%s-%s", p.Name, diag.Name),
				Participant: p,
				WidthWeight: 1.0,
			}).Stage(stage)
			diag.Participant_Shapes = append(diag.Participant_Shapes, ps)
			return ps
		}
		psSales := makePShape(sales)
		psWh := makePShape(warehouse)
		psFin := makePShape(finance)
		_ = psSales
		_ = psWh
		_ = psFin

		// External Participant Shapes
		makeEPShape := func(p *models.Participant, x, y, w, h, tail float64) *models.ExternalParticipantShape {
			eps := (&models.ExternalParticipantShape{
				Name:        fmt.Sprintf("%s-%s", p.Name, diag.Name),
				Participant: p,
				TailHeigth:  tail,
			}).Stage(stage)
			eps.SetX(x)
			eps.SetY(y)
			eps.SetWidth(w)
			eps.SetHeight(h)
			diag.ExternalParticipant_Shapes = append(diag.ExternalParticipant_Shapes, eps)
			return eps
		}
		epsCust := makeEPShape(customer, 40, 150, 160, 65, 840)
		epsCarrier := makeEPShape(carrier, 1370, 480, 170, 65, 480)
		_ = epsCust
		_ = epsCarrier

		// Allocated Resource Shapes
		makeResShape := func(p *models.Participant, r *models.Resource) *models.AllocatedResourceShape {
			ars := (&models.AllocatedResourceShape{
				Name:        fmt.Sprintf("%s-%s-%s", diag.Name, p.Name, r.Name),
				Participant: p,
				Resource:    r,
			}).Stage(stage)
			diag.AllocatedResourceShapes = append(diag.AllocatedResourceShapes, ars)
			return ars
		}
		makeResShape(sales, crmResource)
		makeResShape(warehouse, wmsResource)
		makeResShape(warehouse, erpResource)
		makeResShape(finance, erpResource)
		makeResShape(finance, paymentResource)

		// Task Shapes
		// Sales lane: center around X = 316
		// Wh lane: center around X = 669
		// Fin lane: center around X = 1022
		makeTShape := func(t *models.Task, x, y, w, h float64) *models.TaskShape {
			ts := (&models.TaskShape{
				Name: fmt.Sprintf("%s-%s", t.Name, diag.Name),
				Task: t,
			}).Stage(stage)
			ts.SetX(x)
			ts.SetY(y)
			ts.SetWidth(w)
			ts.SetHeight(h)
			diag.Task_Shapes = append(diag.Task_Shapes, ts)
			return ts
		}

		makeTShape(startSales, 316, 160, 230, 55)
		makeTShape(validateOrder, 316, 260, 230, 65)
		makeTShape(confirmOrder, 316, 380, 230, 65)
		makeTShape(endSales, 316, 500, 230, 55)

		makeTShape(startWh, 669, 160, 230, 55)
		makeTShape(reserveStock, 669, 380, 230, 65)
		makeTShape(pickPack, 669, 500, 230, 65)
		makeTShape(dispatchGoods, 669, 620, 230, 65)
		makeTShape(endWh, 669, 740, 230, 55)

		makeTShape(startFin, 1022, 160, 230, 55)
		makeTShape(issueInvoice, 1022, 620, 230, 65)
		makeTShape(processPayment, 1022, 740, 230, 65)
		makeTShape(closeOrder, 1022, 860, 230, 65)

		// ControlFlow Shapes
		makeCFShape := func(cf *models.ControlFlow) *models.ControlFlowShape {
			cfs := (&models.ControlFlowShape{
				Name:        cf.Name,
				ControlFlow: cf,
				LinkShape: models.LinkShape{
					StartOrientation: models.ORIENTATION_VERTICAL,
					EndOrientation:   models.ORIENTATION_VERTICAL,
					StartRatio:       0.5,
					EndRatio:         0.5,
				},
			}).Stage(stage)
			diag.ControlFlow_Shapes = append(diag.ControlFlow_Shapes, cfs)
			return cfs
		}

		makeCFShape(cfSales1)
		makeCFShape(cfSales2)
		makeCFShape(cfSales3)

		makeCFShape(cfWh1)
		makeCFShape(cfWh2)
		makeCFShape(cfWh3)
		makeCFShape(cfWh4)

		makeCFShape(cfFin1)
		makeCFShape(cfFin2)
		makeCFShape(cfFin3)

		// DataFlow Shapes
		makeDFShape := func(df *models.DataFlow, startOri, endOri models.OrientationType, startRatio, endRatio float64) *models.DataFlowShape {
			dfs := (&models.DataFlowShape{
				Name:     df.Name,
				DataFlow: df,
				LinkShape: models.LinkShape{
					StartOrientation: startOri,
					EndOrientation:   endOri,
					StartRatio:       startRatio,
					EndRatio:         endRatio,
				},
			}).Stage(stage)
			diag.DataFlow_Shapes = append(diag.DataFlow_Shapes, dfs)

			// DataShapes for each data item
			for _, d := range df.Datas {
				ds := (&models.DataShape{
					Name:     fmt.Sprintf("%s-%s-%s", df.Name, df.Name, diag.Name),
					Data:     d,
					DataFlow: df,
				}).Stage(stage)
				diag.Data_Shapes = append(diag.Data_Shapes, ds)
			}
			return dfs
		}

		makeDFShape(df1, models.ORIENTATION_HORIZONTAL, models.ORIENTATION_HORIZONTAL, 0.5, 0.5)
		makeDFShape(df2, models.ORIENTATION_HORIZONTAL, models.ORIENTATION_HORIZONTAL, 0.5, 0.5)
		makeDFShape(df3, models.ORIENTATION_HORIZONTAL, models.ORIENTATION_HORIZONTAL, 0.5, 0.5)
		makeDFShape(df4, models.ORIENTATION_HORIZONTAL, models.ORIENTATION_HORIZONTAL, 0.5, 0.5)
		makeDFShape(df5, models.ORIENTATION_HORIZONTAL, models.ORIENTATION_HORIZONTAL, 0.5, 0.5)
		makeDFShape(df6, models.ORIENTATION_HORIZONTAL, models.ORIENTATION_HORIZONTAL, 0.5, 0.5)
		makeDFShape(df7, models.ORIENTATION_HORIZONTAL, models.ORIENTATION_HORIZONTAL, 0.5, 0.5)

		return diag
	}

	// Diagram 2: Order Fulfillment & Delivery Focus
	createDiagram2 := func() *models.DiagramProcess {
		diag := (&models.DiagramProcess{
			Name:             "Order Fulfillment & Delivery",
			Description:      "Operational focus on sales order processing and warehouse dispatch.",
			IsChecked:        false,
			IsEditable_:      true,
			DefaultBoxWidth:  240,
			DefaultBoxHeigth: 65,
			Width:            1450,
			Height:           950,
		}).Stage(stage)

		pShape := (&models.ProcessShape{
			Name:    "ProcessShape",
			Process: o2cProcess,
		}).Stage(stage)
		pShape.SetX(250)
		pShape.SetY(40)
		pShape.SetWidth(750)
		pShape.SetHeight(850)
		diag.Process_Shapes = append(diag.Process_Shapes, pShape)

		makePShape := func(p *models.Participant) *models.ParticipantShape {
			ps := (&models.ParticipantShape{
				Name:        fmt.Sprintf("%s-%s", p.Name, diag.Name),
				Participant: p,
				WidthWeight: 1.0,
			}).Stage(stage)
			diag.Participant_Shapes = append(diag.Participant_Shapes, ps)
			return ps
		}
		makePShape(sales)
		makePShape(warehouse)

		makeEPShape := func(p *models.Participant, x, y, w, h, tail float64) *models.ExternalParticipantShape {
			eps := (&models.ExternalParticipantShape{
				Name:        fmt.Sprintf("%s-%s", p.Name, diag.Name),
				Participant: p,
				TailHeigth:  tail,
			}).Stage(stage)
			eps.SetX(x)
			eps.SetY(y)
			eps.SetWidth(w)
			eps.SetHeight(h)
			diag.ExternalParticipant_Shapes = append(diag.ExternalParticipant_Shapes, eps)
			return eps
		}
		makeEPShape(customer, 40, 150, 160, 65, 700)
		makeEPShape(carrier, 1060, 450, 170, 65, 400)

		makeTShape := func(t *models.Task, x, y, w, h float64) *models.TaskShape {
			ts := (&models.TaskShape{
				Name: fmt.Sprintf("%s-%s", t.Name, diag.Name),
				Task: t,
			}).Stage(stage)
			ts.SetX(x)
			ts.SetY(y)
			ts.SetWidth(w)
			ts.SetHeight(h)
			diag.Task_Shapes = append(diag.Task_Shapes, ts)
			return ts
		}

		makeTShape(startSales, 316, 160, 230, 55)
		makeTShape(validateOrder, 316, 260, 230, 65)
		makeTShape(confirmOrder, 316, 380, 230, 65)
		makeTShape(endSales, 316, 500, 230, 55)

		makeTShape(startWh, 669, 160, 230, 55)
		makeTShape(reserveStock, 669, 380, 230, 65)
		makeTShape(pickPack, 669, 500, 230, 65)
		makeTShape(dispatchGoods, 669, 620, 230, 65)
		makeTShape(endWh, 669, 740, 230, 55)

		makeCFShape := func(cf *models.ControlFlow) *models.ControlFlowShape {
			cfs := (&models.ControlFlowShape{
				Name:        cf.Name,
				ControlFlow: cf,
				LinkShape: models.LinkShape{
					StartOrientation: models.ORIENTATION_VERTICAL,
					EndOrientation:   models.ORIENTATION_VERTICAL,
					StartRatio:       0.5,
					EndRatio:         0.5,
				},
			}).Stage(stage)
			diag.ControlFlow_Shapes = append(diag.ControlFlow_Shapes, cfs)
			return cfs
		}

		makeCFShape(cfSales1)
		makeCFShape(cfSales2)
		makeCFShape(cfSales3)

		makeCFShape(cfWh1)
		makeCFShape(cfWh2)
		makeCFShape(cfWh3)
		makeCFShape(cfWh4)

		makeDFShape := func(df *models.DataFlow) *models.DataFlowShape {
			dfs := (&models.DataFlowShape{
				Name:     df.Name,
				DataFlow: df,
				LinkShape: models.LinkShape{
					StartOrientation: models.ORIENTATION_HORIZONTAL,
					EndOrientation:   models.ORIENTATION_HORIZONTAL,
					StartRatio:       0.5,
					EndRatio:         0.5,
				},
			}).Stage(stage)
			diag.DataFlow_Shapes = append(diag.DataFlow_Shapes, dfs)

			for _, d := range df.Datas {
				ds := (&models.DataShape{
					Name:     fmt.Sprintf("%s-%s-%s", df.Name, df.Name, diag.Name),
					Data:     d,
					DataFlow: df,
				}).Stage(stage)
				diag.Data_Shapes = append(diag.Data_Shapes, ds)
			}
			return dfs
		}

		makeDFShape(df1)
		makeDFShape(df2)
		makeDFShape(df3)
		makeDFShape(df4)

		return diag
	}

	// Diagram 3: Invoicing & Payment Settlement
	createDiagram3 := func() *models.DiagramProcess {
		diag := (&models.DiagramProcess{
			Name:             "Invoicing & Payment Settlement",
			Description:      "Financial flow between customer billing and accounts receivable reconciliation.",
			IsChecked:        false,
			IsEditable_:      true,
			DefaultBoxWidth:  240,
			DefaultBoxHeigth: 65,
			Width:            1400,
			Height:           850,
		}).Stage(stage)

		pShape := (&models.ProcessShape{
			Name:    "ProcessShape",
			Process: o2cProcess,
		}).Stage(stage)
		pShape.SetX(250)
		pShape.SetY(40)
		pShape.SetWidth(750)
		pShape.SetHeight(750)
		diag.Process_Shapes = append(diag.Process_Shapes, pShape)

		makePShape := func(p *models.Participant) *models.ParticipantShape {
			ps := (&models.ParticipantShape{
				Name:        fmt.Sprintf("%s-%s", p.Name, diag.Name),
				Participant: p,
				WidthWeight: 1.0,
			}).Stage(stage)
			diag.Participant_Shapes = append(diag.Participant_Shapes, ps)
			return ps
		}
		makePShape(sales)
		makePShape(finance)

		makeEPShape := func(p *models.Participant, x, y, w, h, tail float64) *models.ExternalParticipantShape {
			eps := (&models.ExternalParticipantShape{
				Name:        fmt.Sprintf("%s-%s", p.Name, diag.Name),
				Participant: p,
				TailHeigth:  tail,
			}).Stage(stage)
			eps.SetX(x)
			eps.SetY(y)
			eps.SetWidth(w)
			eps.SetHeight(h)
			diag.ExternalParticipant_Shapes = append(diag.ExternalParticipant_Shapes, eps)
			return eps
		}
		makeEPShape(customer, 40, 150, 160, 65, 600)

		makeTShape := func(t *models.Task, x, y, w, h float64) *models.TaskShape {
			ts := (&models.TaskShape{
				Name: fmt.Sprintf("%s-%s", t.Name, diag.Name),
				Task: t,
			}).Stage(stage)
			ts.SetX(x)
			ts.SetY(y)
			ts.SetWidth(w)
			ts.SetHeight(h)
			diag.Task_Shapes = append(diag.Task_Shapes, ts)
			return ts
		}

		makeTShape(startSales, 316, 160, 230, 55)
		makeTShape(validateOrder, 316, 260, 230, 65)
		makeTShape(confirmOrder, 316, 380, 230, 65)
		makeTShape(endSales, 316, 500, 230, 55)

		makeTShape(startFin, 669, 160, 230, 55)
		makeTShape(issueInvoice, 669, 380, 230, 65)
		makeTShape(processPayment, 669, 500, 230, 65)
		makeTShape(closeOrder, 669, 620, 230, 65)

		makeCFShape := func(cf *models.ControlFlow) *models.ControlFlowShape {
			cfs := (&models.ControlFlowShape{
				Name:        cf.Name,
				ControlFlow: cf,
				LinkShape: models.LinkShape{
					StartOrientation: models.ORIENTATION_VERTICAL,
					EndOrientation:   models.ORIENTATION_VERTICAL,
					StartRatio:       0.5,
					EndRatio:         0.5,
				},
			}).Stage(stage)
			diag.ControlFlow_Shapes = append(diag.ControlFlow_Shapes, cfs)
			return cfs
		}

		makeCFShape(cfSales1)
		makeCFShape(cfSales2)
		makeCFShape(cfSales3)

		makeCFShape(cfFin1)
		makeCFShape(cfFin2)
		makeCFShape(cfFin3)

		makeDFShape := func(df *models.DataFlow) *models.DataFlowShape {
			dfs := (&models.DataFlowShape{
				Name:     df.Name,
				DataFlow: df,
				LinkShape: models.LinkShape{
					StartOrientation: models.ORIENTATION_HORIZONTAL,
					EndOrientation:   models.ORIENTATION_HORIZONTAL,
					StartRatio:       0.5,
					EndRatio:         0.5,
				},
			}).Stage(stage)
			diag.DataFlow_Shapes = append(diag.DataFlow_Shapes, dfs)

			for _, d := range df.Datas {
				ds := (&models.DataShape{
					Name:     fmt.Sprintf("%s-%s-%s", df.Name, df.Name, diag.Name),
					Data:     d,
					DataFlow: df,
				}).Stage(stage)
				diag.Data_Shapes = append(diag.Data_Shapes, ds)
			}
			return dfs
		}

		makeDFShape(df1)
		makeDFShape(df2)
		makeDFShape(df6)
		makeDFShape(df7)

		return diag
	}

	diag1 := createDiagram1()
	diag2 := createDiagram2()
	diag3 := createDiagram3()

	o2cProcess.DiagramProcesss = append(o2cProcess.DiagramProcesss, diag1, diag2, diag3)

	// Marshall to order_to_cash.go and stage.go
	stage.MarshallFile("../cmd/process/data/order_to_cash.go", "github.com/fullstack-lang/gong/dsm/process/go/models", "main")
	stage.MarshallFile("../cmd/process/data/stage.go", "github.com/fullstack-lang/gong/dsm/process/go/models", "main")
}

func TestLoadAndVerifyStage(t *testing.T) {
	stage := models.NewStage("test_load")
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "../cmd/process/data/stage.go", nil, parser.ParseComments)
	if err != nil {
		t.Fatalf("failed to parse stage.go: %v", err)
	}

	stage.ParseAstFileFromAst(file, fset, true)
	stage.ComputeReverseMaps()
	stage.ComputeInstancesNb()
	stage.ComputeReferenceAndOrders()

	// Verify instances
	var foundCheckedDiagram *models.DiagramProcess
	diagramCount := 0
	for diag := range *stage.GetInstancesSet[*models.DiagramProcess]() {
		diagramCount++
		if diag.IsChecked {
			foundCheckedDiagram = diag
		}
	}

	if diagramCount != 3 {
		t.Errorf("expected 3 diagrams, got %d", diagramCount)
	}
	if foundCheckedDiagram == nil {
		t.Fatalf("expected a checked diagram, got nil")
	}
	if foundCheckedDiagram.Name != "Order to Cash - Full Process" {
		t.Errorf("expected checked diagram 'Order to Cash - Full Process', got '%s'", foundCheckedDiagram.Name)
	}

	// Verify process
	processes := stage.GetInstancesSorted[*models.Process]()
	if len(processes) != 1 {
		t.Fatalf("expected 1 process, got %d", len(processes))
	}
	if processes[0].Name != "Order to Cash" {
		t.Errorf("expected process 'Order to Cash', got '%s'", processes[0].Name)
	}

	// Verify participants
	participants := stage.GetInstancesSorted[*models.Participant]()
	if len(participants) != 5 { // 3 internal + 2 external
		t.Errorf("expected 5 participants, got %d", len(participants))
	}

	// Verify tasks
	tasks := stage.GetInstancesSorted[*models.Task]()
	if len(tasks) != 13 {
		t.Errorf("expected 13 tasks, got %d", len(tasks))
	}

	// Verify data items
	datas := stage.GetInstancesSorted[*models.Data]()
	if len(datas) != 6 {
		t.Errorf("expected 6 data items, got %d", len(datas))
	}

	// Verify data flows
	dataFlows := stage.GetInstancesSorted[*models.DataFlow]()
	if len(dataFlows) != 7 {
		t.Errorf("expected 7 data flows, got %d", len(dataFlows))
	}
}

