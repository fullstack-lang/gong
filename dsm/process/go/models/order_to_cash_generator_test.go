package models_test

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"strings"
	"testing"

	"github.com/fullstack-lang/gong/dsm/process/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/process/go/models"
	form "github.com/fullstack-lang/gong/lib/form/go/models"
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

var orderToCashResourceIcons = map[string]string{
	"CRM System": `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="#1976D2" stroke-width="1.8">
  <rect x="2" y="4" width="20" height="16" rx="2.5" fill="#1976D2" fill-opacity="0.14"/>
  <circle cx="9" cy="10" r="2.5" fill="#1976D2" fill-opacity="0.25"/>
  <path d="M5 16.5 C5 14.3 6.8 13 9 13 C11.2 13 13 14.3 13 16.5" fill="none"/>
  <line x1="15" y1="9.5" x2="19.5" y2="9.5"/>
  <line x1="15" y1="13" x2="19.5" y2="13"/>
</svg>`,
	"ERP System": `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="#6A1B9A" stroke-width="1.8">
  <ellipse cx="12" cy="5.5" rx="8" ry="3" fill="#6A1B9A" fill-opacity="0.2"/>
  <path d="M4 5.5 V18.5 C4 20.2 7.6 21.5 12 21.5 C16.4 21.5 20 20.2 20 18.5 V5.5" fill="#6A1B9A" fill-opacity="0.1"/>
  <path d="M4 12 C4 13.7 7.6 15 12 15 C16.4 15 20 13.7 20 12" fill="none"/>
</svg>`,
	"WMS (Warehouse Management)": `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="#EF6C00" stroke-width="1.8">
  <polygon points="3,9.5 12,3 21,9.5 21,20.5 3,20.5" fill="#EF6C00" fill-opacity="0.14"/>
  <rect x="6.5" y="14" width="4.5" height="6.5" fill="#EF6C00" fill-opacity="0.25"/>
  <rect x="13" y="14" width="4.5" height="6.5" fill="#EF6C00" fill-opacity="0.25"/>
  <rect x="9.75" y="9" width="4.5" height="5" fill="#EF6C00" fill-opacity="0.25"/>
</svg>`,
	"Payment Gateway": `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="#2E7D32" stroke-width="1.8">
  <rect x="2" y="5" width="20" height="14" rx="2.5" fill="#2E7D32" fill-opacity="0.14"/>
  <line x1="2" y1="9.5" x2="22" y2="9.5"/>
  <rect x="5.5" y="13" width="4" height="2.8" rx="0.6" fill="#2E7D32" fill-opacity="0.3"/>
  <line x1="13" y1="14.5" x2="18.5" y2="14.5"/>
</svg>`,
}

var orderToCashParticipantIcons = map[string]string{
	"Customer": `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="#3949AB" stroke-width="1.8">
  <circle cx="12" cy="8" r="4" fill="#3949AB" fill-opacity="0.18"/>
  <path d="M4.5 20 C4.5 15.8 7.8 13.5 12 13.5 C16.2 13.5 19.5 15.8 19.5 20 Z" fill="#3949AB" fill-opacity="0.14"/>
</svg>`,
	"Logistics Carrier": `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="#00897B" stroke-width="1.8">
  <rect x="1.5" y="5" width="13" height="11" rx="1" fill="#00897B" fill-opacity="0.16"/>
  <polygon points="14.5,8.5 19,8.5 22,12 22,16 14.5,16" fill="#00897B" fill-opacity="0.12"/>
  <circle cx="6.5" cy="17.5" r="2.3" fill="#FFFFFF"/>
  <circle cx="18" cy="17.5" r="2.3" fill="#FFFFFF"/>
</svg>`,
	"Sales & Order Desk": `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="#1565C0" stroke-width="1.8">
  <rect x="3" y="7" width="18" height="13" rx="2" fill="#1565C0" fill-opacity="0.14"/>
  <path d="M8.5 7 V5 C8.5 3.9 9.4 3 10.5 3 H13.5 C14.6 3 15.5 3.9 15.5 5 V7" fill="none"/>
  <line x1="3" y1="12.5" x2="21" y2="12.5"/>
  <rect x="10.2" y="10.8" width="3.6" height="3.4" rx="0.6" fill="#FFFFFF"/>
</svg>`,
	"Warehouse & Logistics": `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="#E65100" stroke-width="1.8">
  <polygon points="12,2.5 21,7.5 21,16.5 12,21.5 3,16.5 3,7.5" fill="#E65100" fill-opacity="0.15"/>
  <polyline points="3,7.5 12,12.5 21,7.5" fill="none"/>
  <line x1="12" y1="12.5" x2="12" y2="21.5"/>
  <line x1="7.5" y1="5" x2="16.5" y2="10"/>
</svg>`,
	"Finance & Billing": `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="#2E7D32" stroke-width="1.8">
  <polygon points="3,9 12,3 21,9" fill="#2E7D32" fill-opacity="0.18"/>
  <line x1="5.5" y1="9" x2="5.5" y2="17.5"/>
  <line x1="10" y1="9" x2="10" y2="17.5"/>
  <line x1="14" y1="9" x2="14" y2="17.5"/>
  <line x1="18.5" y1="9" x2="18.5" y2="17.5"/>
  <rect x="2.5" y="17.5" width="19" height="3" rx="0.6" fill="#2E7D32" fill-opacity="0.22"/>
</svg>`,
}

var orderToCashDataIcons = map[string]string{
	"Purchase Order": `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="#1565C0" stroke-width="1.8">
  <circle cx="9" cy="19.5" r="1.7" fill="#1565C0"/>
  <circle cx="17.5" cy="19.5" r="1.7" fill="#1565C0"/>
  <polygon points="5.5,6 21,6 19,15 7.5,15" fill="#1565C0" fill-opacity="0.15"/>
  <polyline points="2,3.5 4.8,3.5 7.5,15" fill="none"/>
</svg>`,
	"Order Confirmation": `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="#2E7D32" stroke-width="1.8">
  <rect x="4" y="2.5" width="16" height="19" rx="2" fill="#2E7D32" fill-opacity="0.14"/>
  <polyline points="8.5,12.5 11,15 16,9.5" fill="none" stroke-width="2.1"/>
</svg>`,
	"Picking List": `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="#EF6C00" stroke-width="1.8">
  <rect x="4" y="4" width="16" height="17.5" rx="2" fill="#EF6C00" fill-opacity="0.14"/>
  <rect x="8.5" y="2" width="7" height="3.5" rx="1" fill="#FFFFFF"/>
  <line x1="8" y1="10" x2="16.5" y2="10"/>
  <line x1="8" y1="14" x2="16.5" y2="14"/>
  <line x1="8" y1="18" x2="13.5" y2="18"/>
</svg>`,
	"Bill of Lading": `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="#00796B" stroke-width="1.8">
  <polygon points="5,2.5 14.5,2.5 19,7 19,21.5 5,21.5" fill="#00796B" fill-opacity="0.14"/>
  <polyline points="14.5,2.5 14.5,7 19,7" fill="none"/>
  <line x1="8" y1="11.5" x2="16" y2="11.5"/>
  <line x1="8" y1="15.5" x2="16" y2="15.5"/>
</svg>`,
	"Commercial Invoice": `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="#6A1B9A" stroke-width="1.8">
  <polygon points="4.5,2.5 19.5,2.5 19.5,21 17,19.2 14.5,21 12,19.2 9.5,21 7,19.2 4.5,21" fill="#6A1B9A" fill-opacity="0.14"/>
  <line x1="8" y1="7.5" x2="16" y2="7.5"/>
  <line x1="8" y1="11.5" x2="16" y2="11.5"/>
  <line x1="8" y1="15" x2="13" y2="15"/>
</svg>`,
	"Payment Remittance": `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="#1B5E20" stroke-width="1.8">
  <rect x="2" y="5.5" width="20" height="13" rx="2" fill="#1B5E20" fill-opacity="0.15"/>
  <circle cx="12" cy="12" r="3" fill="#1B5E20" fill-opacity="0.24"/>
  <line x1="5.5" y1="9" x2="5.5" y2="15"/>
  <line x1="18.5" y1="9" x2="18.5" y2="15"/>
</svg>`,
}

func TestGenerateOrderToCashModel(t *testing.T) {
	stage := models.NewStage("test")

	// 1. Root Library
	rootLibrary := (&models.Library{
		Name:              "Order to Cash",
		IsRootLibrary:     true,
		NbPixPerCharacter: 8.5,
	}).Stage(stage)

	// 2. Resources
	crmResource := (&models.Resource{Name: "CRM System", SVG: orderToCashResourceIcons["CRM System"]}).Stage(stage)
	erpResource := (&models.Resource{Name: "ERP System", SVG: orderToCashResourceIcons["ERP System"]}).Stage(stage)
	wmsResource := (&models.Resource{Name: "WMS (Warehouse Management)", SVG: orderToCashResourceIcons["WMS (Warehouse Management)"]}).Stage(stage)
	paymentResource := (&models.Resource{Name: "Payment Gateway", SVG: orderToCashResourceIcons["Payment Gateway"]}).Stage(stage)
	rootLibrary.RootResources = append(rootLibrary.RootResources, crmResource, erpResource, wmsResource, paymentResource)

	// 3. Data items
	poData := (&models.Data{Name: "Purchase Order", Description: "Customer Purchase Order with line items and pricing", SVG: orderToCashDataIcons["Purchase Order"]}).Stage(stage)
	ackData := (&models.Data{Name: "Order Confirmation", Description: "Sales order confirmation with promised shipping schedule", SVG: orderToCashDataIcons["Order Confirmation"]}).Stage(stage)
	pickData := (&models.Data{Name: "Picking List", Description: "Itemized warehouse pick and pack instruction slip", SVG: orderToCashDataIcons["Picking List"]}).Stage(stage)
	bolData := (&models.Data{Name: "Bill of Lading", Description: "Freight waybill and carrier consignment tracking record", SVG: orderToCashDataIcons["Bill of Lading"]}).Stage(stage)
	invData := (&models.Data{Name: "Commercial Invoice", Description: "Tax and commercial invoice specifying payment terms", SVG: orderToCashDataIcons["Commercial Invoice"]}).Stage(stage)
	remitData := (&models.Data{Name: "Payment Remittance", Description: "Electronic funds transfer confirmation and remittance advice", SVG: orderToCashDataIcons["Payment Remittance"]}).Stage(stage)
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
		SVG:         orderToCashParticipantIcons["Customer"],
	}).Stage(stage)

	carrier := (&models.Participant{
		Name:        "Logistics Carrier",
		Description: "Third-party freight carrier handling parcel dispatch and physical delivery.",
		SVG:         orderToCashParticipantIcons["Logistics Carrier"],
	}).Stage(stage)

	o2cProcess.ExternalParticipants = append(o2cProcess.ExternalParticipants, customer, carrier)

	// 6. Internal Participants (Swimlanes)
	sales := (&models.Participant{
		Name:        "Sales & Order Desk",
		Description: "Internal sales operations validating customer orders, checking credit, and confirming sales agreements.",
		SVG:         orderToCashParticipantIcons["Sales & Order Desk"],
		Resources:   []*models.Resource{crmResource},
	}).Stage(stage)

	warehouse := (&models.Participant{
		Name:        "Warehouse & Logistics",
		Description: "Fulfillment center checking inventory, picking, packing, and dispatching products to carriers.",
		SVG:         orderToCashParticipantIcons["Warehouse & Logistics"],
		Resources:   []*models.Resource{wmsResource, erpResource},
	}).Stage(stage)

	finance := (&models.Participant{
		Name:        "Finance & Billing",
		Description: "Accounting department generating legal commercial invoices and reconciling customer payment remittances.",
		SVG:         orderToCashParticipantIcons["Finance & Billing"],
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

	// Marshall to order_to_cash.go and stage.go only when explicitly requested
	// (e.g. GENERATE_ORDER_TO_CASH=true) so normal 'go test' runs do not clobber
	// manual diagram edits in data/stage.go.
	if os.Getenv("GENERATE_ORDER_TO_CASH") == "true" {
		stage.MarshallFile("../cmd/process/data/order_to_cash.go", "github.com/fullstack-lang/gong/dsm/process/go/models", "main")
		stage.MarshallFile("../cmd/process/data/stage.go", "github.com/fullstack-lang/gong/dsm/process/go/models", "main")
	}
	if os.Getenv("GENERATE_ORDER_TO_CASH_ONLY") == "true" {
		stage.MarshallFile("../cmd/process/data/order_to_cash.go", "github.com/fullstack-lang/gong/dsm/process/go/models", "main")
	}
	if os.Getenv("UPDATE_STAGE_ICONS") == "true" {
		stageData := models.NewStage("update_stage")
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, "../cmd/process/data/stage.go", nil, parser.ParseComments)
		if err != nil {
			t.Fatalf("failed to parse stage.go: %v", err)
		}
		stageData.ParseAstFileFromAst(file, fset, true)
		stageData.ComputeReverseMaps()
		stageData.ComputeInstancesNb()
		stageData.ComputeReferenceAndOrders()

		for _, r := range stageData.GetInstancesSorted[*models.Resource]() {
			if icon, ok := orderToCashResourceIcons[r.Name]; ok {
				r.SVG = icon
			}
		}
		for _, p := range stageData.GetInstancesSorted[*models.Participant]() {
			if icon, ok := orderToCashParticipantIcons[p.Name]; ok {
				p.SVG = icon
			}
		}
		for _, d := range stageData.GetInstancesSorted[*models.Data]() {
			if icon, ok := orderToCashDataIcons[d.Name]; ok {
				d.SVG = icon
			}
		}
		for _, ds := range stageData.GetInstancesSorted[*models.DataShape]() {
			if ds.Data != nil && ds.Data.SVG != "" && ds.GetWidth() == 180.0 {
				ds.SetWidth(195.0)
			}
		}
		stageData.MarshallFile("../cmd/process/data/stage.go", "github.com/fullstack-lang/gong/dsm/process/go/models", "main")
	}
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

func TestIssue1500_EditSimpleGoAndEnforceSemantic(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("process", "../cmd/process/data/simple.go", "", true, true, true)
	stage := stack.Stage

	// Verify that enum values on DataFlow and DataFlowShape were properly unmarshalled via StageSet
	dataFlows := stage.GetInstancesSorted[*models.DataFlow]()
	if len(dataFlows) != 2 {
		t.Fatalf("expected 2 DataFlows in simple.go, got %d", len(dataFlows))
	}
	for _, df := range dataFlows {
		if df.Type != models.DataFlow_Task2Task && df.Type != models.DataFlow_ExternalParticipant2Task {
			t.Errorf("expected DataFlow %q Type to be unmarshalled, got %q", df.Name, df.Type)
		}
	}

	dataFlowShapes := stage.GetInstancesSorted[*models.DataFlowShape]()
	if len(dataFlowShapes) != 2 {
		t.Fatalf("expected 2 DataFlowShapes in simple.go, got %d", len(dataFlowShapes))
	}
	for _, dfs := range dataFlowShapes {
		if dfs.StartOrientation == "" || dfs.EndOrientation == "" {
			t.Errorf("expected DataFlowShape %q orientations to be unmarshalled, got start=%q end=%q",
				dfs.Name, dfs.StartOrientation, dfs.EndOrientation)
		}
	}

	// Run NewStager (which runs enforceSemantic) exactly as 'go run . edit data/simple.go' does
	models.NewStager(stack.R, stack.Stage, stack.Probe)

	// Verify DataFlows and DataFlowShapes remain staged after enforceSemantic
	if got := len(stage.GetInstancesSorted[*models.DataFlow]()); got != 2 {
		t.Errorf("expected 2 DataFlows after enforceSemantic, got %d", got)
	}
	if got := len(stage.GetInstancesSorted[*models.DataFlowShape]()); got != 2 {
		t.Errorf("expected 2 DataFlowShapes after enforceSemantic, got %d", got)
	}

	// Also verify that if a DataFlow is unstaged (leaving a dangling DataFlowShape),
	// Commit (which runs beforeCommit -> enforceSemantic) cleans and unstages the shape without nil pointer dereference
	dataFlows = stage.GetInstancesSorted[*models.DataFlow]()
	dataFlows[0].UnstageVoid(stage)
	stage.Commit()

	if got := len(stage.GetInstancesSorted[*models.DataFlow]()); got != 1 {
		t.Errorf("expected 1 DataFlow after unstaging one, got %d", got)
	}
	if got := len(stage.GetInstancesSorted[*models.DataFlowShape]()); got != 1 {
		t.Errorf("expected 1 DataFlowShape after unstaging its DataFlow, got %d", got)
	}
}

func TestIssue1503_DataShapeSelectAndResize(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("process", "../cmd/process/data/stage.go", "", true, true, true)
	stager := models.NewStager(stack.R, stack.Stage, stack.Probe)

	svgObj := stager.GetSvgObject()
	if svgObj == nil || len(svgObj.Layers) == 0 {
		t.Fatalf("expected SVG object with at least 1 layer")
	}
	layer := svgObj.Layers[0]

	// Find the svg.Rect for "Order Confirmation"
	var orderConfirmationRect *svg.Rect
	for _, r := range layer.Rects {
		if r.Name == "Order Confirmation" {
			orderConfirmationRect = r
			break
		}
	}
	if orderConfirmationRect == nil {
		t.Fatalf("expected to find an svg.Rect for Data 'Order Confirmation'")
	}
	if !orderConfirmationRect.CanHaveLeftHandle || !orderConfirmationRect.CanHaveRightHandle {
		t.Errorf("expected Data svg.Rect to have left and right resize handles enabled")
	}
	if orderConfirmationRect.OnSelect == nil || orderConfirmationRect.OnResize == nil {
		t.Fatalf("expected Data svg.Rect to have OnSelect and OnResize callbacks")
	}

	// 1. Selecting the Data rect should update the Probe Form with the Data ("Order Confirmation"), not the Participant
	orderConfirmationRect.OnSelect()
	formStage := stack.Probe.GetFormStage()
	formGroups := formStage.GetInstancesSorted[*form.FormGroup]()
	if len(formGroups) == 0 {
		t.Fatalf("expected Probe FormGroup to be populated after selecting Data rect")
	}
	if formGroups[0].Label != "Order Confirmation" {
		t.Errorf("expected Probe FormGroup Label to be 'Order Confirmation', got %q", formGroups[0].Label)
	}

	// 2. Resizing the Data rect narrower should wrap the text into multiple lines and increase rect height
	initialHeight := orderConfirmationRect.Height
	orderConfirmationRect.OnResize(orderConfirmationRect.X, orderConfirmationRect.Y, 90.0, orderConfirmationRect.Height)

	svgObjAfter := stager.GetSvgObject()
	var resizedRect *svg.Rect
	for _, r := range svgObjAfter.Layers[0].Rects {
		if r.Name == "Order Confirmation" && r.Width == 90.0 {
			resizedRect = r
			break
		}
	}
	if resizedRect == nil {
		t.Fatalf("expected to find resized svg.Rect with Width=90.0 for 'Order Confirmation'")
	}
	if len(resizedRect.RectAnchoredTexts) == 0 {
		t.Fatalf("expected RectAnchoredText inside Data svg.Rect")
	}
	wrappedContent := resizedRect.RectAnchoredTexts[0].Content
	if !strings.Contains(wrappedContent, "\n") {
		t.Errorf("expected 'Order Confirmation' to wrap across multiple lines at Width=90.0, got %q", wrappedContent)
	}
	if resizedRect.Height <= initialHeight {
		t.Errorf("expected wrapped Data rect height (%v) to be greater than single-line height (%v)", resizedRect.Height, initialHeight)
	}
}

func TestIssue1528_ParticipantResourceAndDataSVGIcons(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("process", "../cmd/process/data/stage.go", "", true, true, true)
	stager := models.NewStager(stack.R, stack.Stage, stack.Probe)

	svgObj := stager.GetSvgObject()
	if svgObj == nil || len(svgObj.Layers) == 0 {
		t.Fatalf("expected SVG object with at least 1 layer")
	}
	layer := svgObj.Layers[0]

	rectByName := make(map[string]*svg.Rect)
	for _, r := range layer.Rects {
		rectByName[r.Name] = r
	}

	// 1. Internal swimlane participants: header icon (negative Y_Offset) + allocated resource icons (positive Y_Offset)
	for _, laneName := range []string{
		"Sales & Order Desk-Order to Cash - Full Process",
		"Warehouse & Logistics-Order to Cash - Full Process",
		"Finance & Billing-Order to Cash - Full Process",
	} {
		r := rectByName[laneName]
		if r == nil {
			t.Fatalf("missing swimlane rect %q", laneName)
		}
		hasHeaderIcon := false
		hasResourceIcon := false
		for _, p := range r.RectAnchoredPaths {
			if p.Y_Offset < 0 {
				hasHeaderIcon = true
			}
			if p.Y_Offset > 0 {
				hasResourceIcon = true
			}
		}
		if !hasHeaderIcon {
			t.Errorf("swimlane %q missing participant header icon RectAnchoredPath", laneName)
		}
		if !hasResourceIcon {
			t.Errorf("swimlane %q missing allocated resource icon RectAnchoredPath", laneName)
		}
		if len(r.RectAnchoredTexts) == 0 || r.RectAnchoredTexts[0].X_Offset <= 0 {
			t.Errorf("swimlane %q title was not shifted right for left icon", laneName)
		}
	}

	// 2. External participants: left-aligned icon + shifted title
	for _, extName := range []string{
		"Customer-Order to Cash - Full Process",
		"Logistics Carrier-Order to Cash - Full Process",
	} {
		r := rectByName[extName]
		if r == nil {
			t.Fatalf("missing external participant rect %q", extName)
		}
		if len(r.RectAnchoredPaths) == 0 {
			t.Errorf("external participant %q missing icon RectAnchoredPaths", extName)
		}
		if len(r.RectAnchoredTexts) == 0 || r.RectAnchoredTexts[0].X_Offset <= 0 {
			t.Errorf("external participant %q title was not shifted right for left icon", extName)
		}
	}

	// 3. Data items on data flows: left-aligned 16x16 icon + leftMargin=24
	for _, dataName := range []string{
		"Purchase Order",
		"Order Confirmation",
		"Bill of Lading",
		"Commercial Invoice",
		"Payment Remittance",
	} {
		r := rectByName[dataName]
		if r == nil {
			t.Fatalf("missing data rect %q", dataName)
		}
		if len(r.RectAnchoredPaths) == 0 {
			t.Errorf("data rect %q missing icon RectAnchoredPaths", dataName)
		}
		if len(r.RectAnchoredTexts) == 0 || r.RectAnchoredTexts[0].X_Offset != 24.0 {
			t.Errorf("data rect %q expected title X_Offset=24.0, got %v", dataName, r.RectAnchoredTexts[0].X_Offset)
		}
	}

	// 4. Static website SVG string export includes rendered icon paths
	svgOut, _, _ := svgObj.GenerateString()
	if !strings.Contains(svgOut, "<svg") || !strings.Contains(svgOut, "scale(") {
		t.Errorf("expected svgObj.GenerateString() to output valid SVG with scaled icon paths")
	}
}


