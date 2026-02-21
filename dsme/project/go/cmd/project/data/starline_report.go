package main

import (
	"time"

	"github.com/fullstack-lang/gong/dsme/project/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage

	__Diagram__00000000_ := (&models.Diagram{Name: `Default Diagram`}).Stage(stage)
	__Diagram__00000001_ := (&models.Diagram{Name: `WBS`}).Stage(stage)
	__Diagram__00000002_ := (&models.Diagram{Name: `PBS`}).Stage(stage)
	__Diagram__00000003_ := (&models.Diagram{Name: `RBS`}).Stage(stage)
	__Diagram__00000004_ := (&models.Diagram{Name: `PIT focus`}).Stage(stage)
	__Diagram__00000005_ := (&models.Diagram{Name: `PIT Report`}).Stage(stage)
	__Diagram__00000006_ := (&models.Diagram{Name: `RCS PBS`}).Stage(stage)

	__Note__00000000_ := (&models.Note{Name: `CFT ended in march 2025`}).Stage(stage)
	__Note__00000001_ := (&models.Note{Name: `A thorough review of the STAR report is advised (p14)
`}).Stage(stage)
	__Note__00000002_ := (&models.Note{Name: `NASA utilized a firm fixed price contracting type for CCtCap. This was a significant shift from the cost-plus contracting for traditional NASA builds of developmental vehicles. These shifts signified that CCP was not only positioned to be an innovative, first-of-its kind program for NASA, but how it interacted with new and traditional space flight industry providers was setup to be significantly distinct and different.  `}).Stage(stage)
	__Note__00000003_ := (&models.Note{Name: `The CCP 1100 series of requirements were deliberately written at a higher-level, leaving room for provider innovation but there was also room for incorrect/inadequate interpretation by the providers. `}).Stage(stage)
	__Note__00000004_ := (&models.Note{Name: `The Commercial Provider focused on meeting contractual requirement language resulting in insufficient demonstration across the components/system and ground/flight. `}).Stage(stage)
	__Note__00000005_ := (&models.Note{Name: `Suppliers’ build quality/variability issues can be hard to exonerate for service modules, which is hardware that is disposed of during re-entry. `}).Stage(stage)
	__Note__00000006_ := (&models.Note{Name: `The Aerojet Rocketdyne (AR) thermal model included the effects of jet firings, but these effects were not validated by ground testing.

Boeing thermal model did not include the effects of jet firings before CFT. `}).Stage(stage)
	__Note__00000007_ := (&models.Note{Name: `The thruster performance from OFT1 & OFT2 experienced greater than expected temperatures and continuing to operate lead to a false sense of security of the thruster capability/performance. `}).Stage(stage)
	__Note__00000008_ := (&models.Note{Name: `Flight instrumentation locations for thermal sensors were limited and in different locations than the locations for RCS Thruster ground firings. `}).Stage(stage)
	__Note__00000009_ := (&models.Note{Name: `OFT1 & OFT2 investigations did not include RCS/OMAC thruster firings and fault trees were not validated through subsequent ground testing. `}).Stage(stage)
	__Note__00000010_ := (&models.Note{Name: `For OFT2, NASA/Boeing did not have tools to measure thruster degradation, simply treated the thruster as failed/operational. `}).Stage(stage)
	__Note__00000011_ := (&models.Note{Name: `Pc (Chamber Pressure).To know if a thruster is actually firing, the spacecraft's computer looks at the pressure sensor inside the combustion chamber (the Pc telemetry). If the pressure shoots up to the expected level, it means combustion is happening and the thruster is pushing. If the pressure stays low or at zero, it assumes the thruster failed.`}).Stage(stage)
	__Note__00000012_ := (&models.Note{Name: `Pc transducer failed but the engine was able to restart. Boing confidence in the harxware robustness`}).Stage(stage)

	__NoteProductShape__00000000_ := (&models.NoteProductShape{Name: `A thorough review of the STAR report is advised to  Starliner Tests and Anomalies Review (STAR) Investigation Report`}).Stage(stage)
	__NoteProductShape__00000001_ := (&models.NoteProductShape{Name: `NASA utilized a firm fixed price contracting type for CCtCap. This was a significant shift from the cost-plus contracting for traditional NASA builds of developmental vehicles. These shifts signified that CCP was not only positioned to be an innovative, first-of-its kind program for NASA, but how it interacted with new and traditional space flight industry providers was setup to be significantly distinct and different.   to Commercial Crew Transportation Capability (CCtCap).`}).Stage(stage)
	__NoteProductShape__00000002_ := (&models.NoteProductShape{Name: ` to CCP Requirements`}).Stage(stage)
	__NoteProductShape__00000003_ := (&models.NoteProductShape{Name: `The Commercial Provider focused on meeting contractual requirement language resulting in insufficient demonstration across the components/system and ground/flight.  to CCP Requirements`}).Stage(stage)
	__NoteProductShape__00000004_ := (&models.NoteProductShape{Name: `Suppliers’ build quality/variability issues can be hard to exonerate for service modules, which is hardware that is disposed of during re-entry.  to CM RCS Thrusters`}).Stage(stage)
	__NoteProductShape__00000005_ := (&models.NoteProductShape{Name: `Suppliers’ build quality/variability issues can be hard to exonerate for service modules, which is hardware that is disposed of during re-entry.  to SM RCS Thrusters`}).Stage(stage)
	__NoteProductShape__00000006_ := (&models.NoteProductShape{Name: `The Aerojet Rocketdyne (AR) thermal model included the effects of jet firings, but these effects were not validated by ground testing.

Boeing thermal model did not include the effects of jet firings before CFT.  to 28 SM RCS Thrusters`}).Stage(stage)
	__NoteProductShape__00000007_ := (&models.NoteProductShape{Name: `The thruster performance from OFT1 & OFT2 experienced greater than expected temperatures and continuing to operate lead to a false sense of security of the thruster capability/performance.  to 28 SM RCS Thrusters`}).Stage(stage)
	__NoteProductShape__00000008_ := (&models.NoteProductShape{Name: `Flight instrumentation locations for thermal sensors were limited and in different locations than the locations for RCS Thruster ground firings.  to Thermal Sensors`}).Stage(stage)
	__NoteProductShape__00000009_ := (&models.NoteProductShape{Name: `Flight instrumentation locations for thermal sensors were limited and in different locations than the locations for RCS Thruster ground firings.  to 28 SM RCS Thrusters`}).Stage(stage)
	__NoteProductShape__00000010_ := (&models.NoteProductShape{Name: `Flight instrumentation locations for thermal sensors were limited and in different locations than the locations for RCS Thruster ground firings.  to 12 CM RCS Thrusters`}).Stage(stage)
	__NoteProductShape__00000011_ := (&models.NoteProductShape{Name: `OFT1 & OFT2 investigations did not include RCS/OMAC thruster firings and fault trees were not validated through subsequent ground testing.  to 28 SM RCS Thrusters (2800 lbf)`}).Stage(stage)
	__NoteProductShape__00000012_ := (&models.NoteProductShape{Name: `For OFT2, NASA/Boeing did not have tools to measure thruster degradation, simply treated the thruster as failed/operational.  to 28 SM RCS Thrusters (2800 lbf)`}).Stage(stage)
	__NoteProductShape__00000013_ := (&models.NoteProductShape{Name: `Pc (Chamber Pressure) to 12 CM RCS Thrusters (1200 lbf)`}).Stage(stage)
	__NoteProductShape__00000014_ := (&models.NoteProductShape{Name: `Pc (Chamber Pressure).To know if a thruster is actually firing, the spacecraft's computer looks at the pressure sensor inside the combustion chamber (the Pc telemetry). If the pressure shoots up to the expected level, it means combustion is happening and the thruster is pushing. If the pressure stays low or at zero, it assumes the thruster failed. to Transducer / pressure sensor`}).Stage(stage)
	__NoteProductShape__00000015_ := (&models.NoteProductShape{Name: `Pc transducer failed but the engine was able to restart. Boing confidence in the harxware robustness to Transducer / pressure sensor (Pc)`}).Stage(stage)

	__NoteShape__00000000_ := (&models.NoteShape{Name: `-Default Diagram`}).Stage(stage)
	__NoteShape__00000001_ := (&models.NoteShape{Name: `-PIT focus`}).Stage(stage)
	__NoteShape__00000002_ := (&models.NoteShape{Name: `-PBS`}).Stage(stage)
	__NoteShape__00000003_ := (&models.NoteShape{Name: `-PIT Report`}).Stage(stage)
	__NoteShape__00000004_ := (&models.NoteShape{Name: `-PIT Report`}).Stage(stage)
	__NoteShape__00000005_ := (&models.NoteShape{Name: `-RCS PBS`}).Stage(stage)
	__NoteShape__00000006_ := (&models.NoteShape{Name: `-RCS PBS`}).Stage(stage)
	__NoteShape__00000007_ := (&models.NoteShape{Name: `-RCS PBS`}).Stage(stage)
	__NoteShape__00000008_ := (&models.NoteShape{Name: `-RCS PBS`}).Stage(stage)
	__NoteShape__00000009_ := (&models.NoteShape{Name: `-RCS PBS`}).Stage(stage)
	__NoteShape__00000010_ := (&models.NoteShape{Name: `-RCS PBS`}).Stage(stage)
	__NoteShape__00000011_ := (&models.NoteShape{Name: `-RCS PBS`}).Stage(stage)
	__NoteShape__00000012_ := (&models.NoteShape{Name: `-RCS PBS`}).Stage(stage)

	__NoteTaskShape__00000000_ := (&models.NoteTaskShape{Name: `CFT ended in march 2025 to Starliner Crewed Flight Test (CFT)`}).Stage(stage)

	__Product__00000001_ := (&models.Product{Name: `Dragon`}).Stage(stage)
	__Product__00000002_ := (&models.Product{Name: `CST-100 Starliner`}).Stage(stage)
	__Product__00000003_ := (&models.Product{Name: `Reports`}).Stage(stage)
	__Product__00000004_ := (&models.Product{Name: ` Starliner Tests and Anomalies Review (STAR) Investigation Report`}).Stage(stage)
	__Product__00000005_ := (&models.Product{Name: `Program Investigation Team (PIT) Report`}).Stage(stage)
	__Product__00000006_ := (&models.Product{Name: `Commercial Crew Transportation Capability (CCtCap).`}).Stage(stage)
	__Product__00000007_ := (&models.Product{Name: `NASA Assets/Capabities`}).Stage(stage)
	__Product__00000008_ := (&models.Product{Name: `ISS`}).Stage(stage)
	__Product__00000009_ := (&models.Product{Name: `3 Commercial Crew Program (CCP) Background`}).Stage(stage)
	__Product__00000010_ := (&models.Product{Name: `CCP Requirements`}).Stage(stage)
	__Product__00000011_ := (&models.Product{Name: `Crew Module (CM)`}).Stage(stage)
	__Product__00000012_ := (&models.Product{Name: `CM Reaction Control System (RCS)`}).Stage(stage)
	__Product__00000013_ := (&models.Product{Name: `12 CM RCS Thrusters (1200 lbf)`}).Stage(stage)
	__Product__00000014_ := (&models.Product{Name: `Service Module (SM)`}).Stage(stage)
	__Product__00000015_ := (&models.Product{Name: `SM Reaction Control System (SM RCS)`}).Stage(stage)
	__Product__00000016_ := (&models.Product{Name: `28 SM RCS Thrusters (2800 lbf)`}).Stage(stage)
	__Product__00000017_ := (&models.Product{Name: `2 Additional Investigations `}).Stage(stage)
	__Product__00000018_ := (&models.Product{Name: `STAR Report Summary and Findings `}).Stage(stage)
	__Product__00000019_ := (&models.Product{Name: `Thermal Sensors`}).Stage(stage)
	__Product__00000020_ := (&models.Product{Name: `Dog House`}).Stage(stage)
	__Product__00000021_ := (&models.Product{Name: `Orbital Maneuvering and Attitude Control (OMAC), 30 000 lbf`}).Stage(stage)
	__Product__00000022_ := (&models.Product{Name: `Boeing’s Enterprise Root Cause/Corrective Actions (eRCCA)`}).Stage(stage)
	__Product__00000023_ := (&models.Product{Name: `3.1   Orbital Flight Test (OFT) Summary`}).Stage(stage)
	__Product__00000024_ := (&models.Product{Name: `3.2   Orbital Flight Test 2 (OFT-2) Summary`}).Stage(stage)
	__Product__00000025_ := (&models.Product{Name: `Flight Software (FSW)`}).Stage(stage)
	__Product__00000026_ := (&models.Product{Name: `Fault Detection, Isolation, and Recovery (FDIR)`}).Stage(stage)
	__Product__00000027_ := (&models.Product{Name: `3.3 Comparing SM RCS Thrusters Triggering Fail-Off FDIR on OFT1/OFT2`}).Stage(stage)
	__Product__00000028_ := (&models.Product{Name: `Transducer / pressure sensor (Pc)`}).Stage(stage)
	__Product__00000029_ := (&models.Product{Name: `4     Technical Root Cause Analysis (RCA) and Findings`}).Stage(stage)

	__ProductCompositionShape__00000000_ := (&models.ProductCompositionShape{Name: `Reports to `}).Stage(stage)
	__ProductCompositionShape__00000001_ := (&models.ProductCompositionShape{Name: `Reports to `}).Stage(stage)
	__ProductCompositionShape__00000002_ := (&models.ProductCompositionShape{Name: `Commercial Crew Transportation Capability (CCtCap). to Starliner`}).Stage(stage)
	__ProductCompositionShape__00000003_ := (&models.ProductCompositionShape{Name: `Commercial Crew Transportation Capability (CCtCap). to Dragon`}).Stage(stage)
	__ProductCompositionShape__00000004_ := (&models.ProductCompositionShape{Name: `NASA Assets/Capabities to Commercial Crew Transportation Capability (CCtCap).`}).Stage(stage)
	__ProductCompositionShape__00000005_ := (&models.ProductCompositionShape{Name: `Commercial Crew Transportation Capability (CCtCap). to `}).Stage(stage)
	__ProductCompositionShape__00000006_ := (&models.ProductCompositionShape{Name: `Program Investigation Team (PIT) Report to `}).Stage(stage)
	__ProductCompositionShape__00000007_ := (&models.ProductCompositionShape{Name: `CM to `}).Stage(stage)
	__ProductCompositionShape__00000008_ := (&models.ProductCompositionShape{Name: `RCS to `}).Stage(stage)
	__ProductCompositionShape__00000009_ := (&models.ProductCompositionShape{Name: `CST-100 Starliner to CM`}).Stage(stage)
	__ProductCompositionShape__00000010_ := (&models.ProductCompositionShape{Name: `CST-100 Starliner to `}).Stage(stage)
	__ProductCompositionShape__00000011_ := (&models.ProductCompositionShape{Name: `Service Module (SM) to `}).Stage(stage)
	__ProductCompositionShape__00000012_ := (&models.ProductCompositionShape{Name: `SM Reaction Control System (SM RCS) to `}).Stage(stage)
	__ProductCompositionShape__00000013_ := (&models.ProductCompositionShape{Name: `Program Investigation Team (PIT) Report to `}).Stage(stage)
	__ProductCompositionShape__00000014_ := (&models.ProductCompositionShape{Name: `Chap 2 Additional Investigations  to `}).Stage(stage)
	__ProductCompositionShape__00000015_ := (&models.ProductCompositionShape{Name: `Crew Module (CM) to `}).Stage(stage)
	__ProductCompositionShape__00000016_ := (&models.ProductCompositionShape{Name: `12 CM RCS Thrusters to `}).Stage(stage)
	__ProductCompositionShape__00000017_ := (&models.ProductCompositionShape{Name: `CST-100 Starliner to `}).Stage(stage)
	__ProductCompositionShape__00000018_ := (&models.ProductCompositionShape{Name: `Reports to Boeing’s Enterprise Root Cause/Corrective Actions (eRCCA)`}).Stage(stage)
	__ProductCompositionShape__00000019_ := (&models.ProductCompositionShape{Name: `Chap3 Commercial Crew Program (CCP) Background to `}).Stage(stage)
	__ProductCompositionShape__00000020_ := (&models.ProductCompositionShape{Name: `Chap3 Commercial Crew Program (CCP) Background to `}).Stage(stage)
	__ProductCompositionShape__00000021_ := (&models.ProductCompositionShape{Name: `CST-100 Starliner to `}).Stage(stage)
	__ProductCompositionShape__00000022_ := (&models.ProductCompositionShape{Name: `CST-100 Starliner to `}).Stage(stage)
	__ProductCompositionShape__00000023_ := (&models.ProductCompositionShape{Name: `Chap3 Commercial Crew Program (CCP) Background to `}).Stage(stage)
	__ProductCompositionShape__00000024_ := (&models.ProductCompositionShape{Name: `12 CM RCS Thrusters (1200 lbf) to `}).Stage(stage)
	__ProductCompositionShape__00000025_ := (&models.ProductCompositionShape{Name: `Chap3 Commercial Crew Program (CCP) Background to `}).Stage(stage)

	__ProductShape__00000002_ := (&models.ProductShape{Name: `Dragon-PBS`}).Stage(stage)
	__ProductShape__00000003_ := (&models.ProductShape{Name: `Starliner-PBS`}).Stage(stage)
	__ProductShape__00000004_ := (&models.ProductShape{Name: `-PBS`}).Stage(stage)
	__ProductShape__00000005_ := (&models.ProductShape{Name: `-PBS`}).Stage(stage)
	__ProductShape__00000006_ := (&models.ProductShape{Name: `-PBS`}).Stage(stage)
	__ProductShape__00000007_ := (&models.ProductShape{Name: `Program Investigation Team (PIT) Report-PIT focus`}).Stage(stage)
	__ProductShape__00000008_ := (&models.ProductShape{Name: ` Starliner Tests and Anomalies Review (STAR) Investigation Report-PIT focus`}).Stage(stage)
	__ProductShape__00000009_ := (&models.ProductShape{Name: `-PBS`}).Stage(stage)
	__ProductShape__00000010_ := (&models.ProductShape{Name: `-PBS`}).Stage(stage)
	__ProductShape__00000011_ := (&models.ProductShape{Name: `-PBS`}).Stage(stage)
	__ProductShape__00000012_ := (&models.ProductShape{Name: `Program Investigation Team (PIT) Report-PIT Report`}).Stage(stage)
	__ProductShape__00000013_ := (&models.ProductShape{Name: `-PIT Report`}).Stage(stage)
	__ProductShape__00000014_ := (&models.ProductShape{Name: `-PIT Report`}).Stage(stage)
	__ProductShape__00000015_ := (&models.ProductShape{Name: `-RCS PBS`}).Stage(stage)
	__ProductShape__00000016_ := (&models.ProductShape{Name: `-RCS PBS`}).Stage(stage)
	__ProductShape__00000017_ := (&models.ProductShape{Name: `-RCS PBS`}).Stage(stage)
	__ProductShape__00000018_ := (&models.ProductShape{Name: `CST-100 Starliner-RCS PBS`}).Stage(stage)
	__ProductShape__00000019_ := (&models.ProductShape{Name: `-RCS PBS`}).Stage(stage)
	__ProductShape__00000020_ := (&models.ProductShape{Name: `-RCS PBS`}).Stage(stage)
	__ProductShape__00000021_ := (&models.ProductShape{Name: `-RCS PBS`}).Stage(stage)
	__ProductShape__00000022_ := (&models.ProductShape{Name: `-PIT Report`}).Stage(stage)
	__ProductShape__00000023_ := (&models.ProductShape{Name: `-PIT Report`}).Stage(stage)
	__ProductShape__00000024_ := (&models.ProductShape{Name: `-RCS PBS`}).Stage(stage)
	__ProductShape__00000025_ := (&models.ProductShape{Name: `-RCS PBS`}).Stage(stage)
	__ProductShape__00000026_ := (&models.ProductShape{Name: `-RCS PBS`}).Stage(stage)
	__ProductShape__00000027_ := (&models.ProductShape{Name: `Boeing’s Enterprise Root Cause/Corrective Actions (eRCCA)-PBS`}).Stage(stage)
	__ProductShape__00000028_ := (&models.ProductShape{Name: `-PIT Report`}).Stage(stage)
	__ProductShape__00000029_ := (&models.ProductShape{Name: `-PIT Report`}).Stage(stage)
	__ProductShape__00000030_ := (&models.ProductShape{Name: `-RCS PBS`}).Stage(stage)
	__ProductShape__00000031_ := (&models.ProductShape{Name: `-RCS PBS`}).Stage(stage)
	__ProductShape__00000032_ := (&models.ProductShape{Name: `-PIT Report`}).Stage(stage)
	__ProductShape__00000033_ := (&models.ProductShape{Name: `-RCS PBS`}).Stage(stage)
	__ProductShape__00000034_ := (&models.ProductShape{Name: `-PIT Report`}).Stage(stage)

	__Project__00000000_ := (&models.Project{Name: `Startliner Mishape Report`}).Stage(stage)

	__Resource__00000000_ := (&models.Resource{Name: `Program Investigation Team (PIT)`}).Stage(stage)
	__Resource__00000001_ := (&models.Resource{Name: `Barry "Butch" Wilmore`}).Stage(stage)
	__Resource__00000002_ := (&models.Resource{Name: `Sunita "Suni" Williams`}).Stage(stage)
	__Resource__00000003_ := (&models.Resource{Name: `NASA`}).Stage(stage)
	__Resource__00000004_ := (&models.Resource{Name: `Crew Commercial Program (CPP)`}).Stage(stage)
	__Resource__00000005_ := (&models.Resource{Name: `Crews`}).Stage(stage)
	__Resource__00000006_ := (&models.Resource{Name: `Boeing`}).Stage(stage)
	__Resource__00000007_ := (&models.Resource{Name: ``}).Stage(stage)
	__Resource__00000008_ := (&models.Resource{Name: ``}).Stage(stage)
	__Resource__00000009_ := (&models.Resource{Name: ` Starliner Tests and Anomalies Review (STAR) Investigation Team`}).Stage(stage)

	__ResourceCompositionShape__00000000_ := (&models.ResourceCompositionShape{Name: `NASA to `}).Stage(stage)
	__ResourceCompositionShape__00000001_ := (&models.ResourceCompositionShape{Name: `Crew Commercial Program (CPP) to `}).Stage(stage)
	__ResourceCompositionShape__00000002_ := (&models.ResourceCompositionShape{Name: `Crews to Barry "Butch" Wilmore`}).Stage(stage)
	__ResourceCompositionShape__00000003_ := (&models.ResourceCompositionShape{Name: `Crews to Sunita "Suni" Williams`}).Stage(stage)
	__ResourceCompositionShape__00000004_ := (&models.ResourceCompositionShape{Name: `NASA to PITProgram Investigation Team (PIT)`}).Stage(stage)
	__ResourceCompositionShape__00000005_ := (&models.ResourceCompositionShape{Name: `Boeing to `}).Stage(stage)
	__ResourceCompositionShape__00000006_ := (&models.ResourceCompositionShape{Name: `NASA to `}).Stage(stage)
	__ResourceCompositionShape__00000007_ := (&models.ResourceCompositionShape{Name: `Crew Commercial Program (CPP) to `}).Stage(stage)

	__ResourceShape__00000000_ := (&models.ResourceShape{Name: `-Default Diagram`}).Stage(stage)
	__ResourceShape__00000001_ := (&models.ResourceShape{Name: `-Default Diagram`}).Stage(stage)
	__ResourceShape__00000002_ := (&models.ResourceShape{Name: `-Default Diagram`}).Stage(stage)
	__ResourceShape__00000003_ := (&models.ResourceShape{Name: `-RBS`}).Stage(stage)
	__ResourceShape__00000004_ := (&models.ResourceShape{Name: `-RBS`}).Stage(stage)
	__ResourceShape__00000005_ := (&models.ResourceShape{Name: `-RBS`}).Stage(stage)
	__ResourceShape__00000006_ := (&models.ResourceShape{Name: `Barry "Butch" Wilmore-RBS`}).Stage(stage)
	__ResourceShape__00000007_ := (&models.ResourceShape{Name: `Sunita "Suni" Williams-RBS`}).Stage(stage)
	__ResourceShape__00000008_ := (&models.ResourceShape{Name: `PITProgram Investigation Team (PIT)-RBS`}).Stage(stage)
	__ResourceShape__00000009_ := (&models.ResourceShape{Name: `-RBS`}).Stage(stage)
	__ResourceShape__00000010_ := (&models.ResourceShape{Name: `-RBS`}).Stage(stage)
	__ResourceShape__00000011_ := (&models.ResourceShape{Name: `-RBS`}).Stage(stage)
	__ResourceShape__00000012_ := (&models.ResourceShape{Name: `-RBS`}).Stage(stage)
	__ResourceShape__00000013_ := (&models.ResourceShape{Name: `Program Investigation Team (PIT)-PBS`}).Stage(stage)
	__ResourceShape__00000014_ := (&models.ResourceShape{Name: `Program Investigation Team (PIT)-PIT focus`}).Stage(stage)

	__ResourceTaskShape__00000000_ := (&models.ResourceTaskShape{Name: `PITProgram Investigation Team (PIT) to Mishap investigation`}).Stage(stage)

	__Root__00000000_ := (&models.Root{Name: `Root`}).Stage(stage)

	__Task__00000000_ := (&models.Task{Name: `Starliner Crewed Flight Test (CFT)`}).Stage(stage)
	__Task__00000001_ := (&models.Task{Name: `Mishap investigations`}).Stage(stage)
	__Task__00000002_ := (&models.Task{Name: `Commercial Crew Program (CCP),`}).Stage(stage)
	__Task__00000003_ := (&models.Task{Name: ` Commercial ReSupply (CRS) `}).Stage(stage)
	__Task__00000004_ := (&models.Task{Name: `Program Investigation Team (PIT) Report`}).Stage(stage)
	__Task__00000005_ := (&models.Task{Name: `Orbital Flight Test-1 (OFT-1)`}).Stage(stage)
	__Task__00000006_ := (&models.Task{Name: `Orbital Flight Test-2 (OFT-2)`}).Stage(stage)

	__TaskCompositionShape__00000000_ := (&models.TaskCompositionShape{Name: `Commercial Crew Program (CCP), to Starliner Crewed Flight Test (CFT)`}).Stage(stage)
	__TaskCompositionShape__00000001_ := (&models.TaskCompositionShape{Name: `Commercial Crew Program (CCP), to Starliner Crewed Flight Test (CFT)`}).Stage(stage)
	__TaskCompositionShape__00000002_ := (&models.TaskCompositionShape{Name: `Mishap investigations to `}).Stage(stage)
	__TaskCompositionShape__00000003_ := (&models.TaskCompositionShape{Name: `Commercial Crew Program (CCP), to `}).Stage(stage)
	__TaskCompositionShape__00000004_ := (&models.TaskCompositionShape{Name: `Commercial Crew Program (CCP), to `}).Stage(stage)

	__TaskInputShape__00000000_ := (&models.TaskInputShape{Name: `Program Investigation Team (PIT) Report to  Starliner Tests and Anomalies Review (STAR) Investigation Report`}).Stage(stage)

	__TaskOutputShape__00000000_ := (&models.TaskOutputShape{Name: `Program Investigation Team (PIT) Report to Program Investigation Team (PIT) Report`}).Stage(stage)

	__TaskShape__00000000_ := (&models.TaskShape{Name: `NewTask-Default Diagram`}).Stage(stage)
	__TaskShape__00000001_ := (&models.TaskShape{Name: `NewTask-Default Diagram`}).Stage(stage)
	__TaskShape__00000002_ := (&models.TaskShape{Name: `-Default Diagram`}).Stage(stage)
	__TaskShape__00000003_ := (&models.TaskShape{Name: `Mishap investigation-WBS`}).Stage(stage)
	__TaskShape__00000004_ := (&models.TaskShape{Name: `Commercial Crew Program (CCP),-WBS`}).Stage(stage)
	__TaskShape__00000005_ := (&models.TaskShape{Name: `Starliner Crewed Flight Test (CFT)-WBS`}).Stage(stage)
	__TaskShape__00000006_ := (&models.TaskShape{Name: `-WBS`}).Stage(stage)
	__TaskShape__00000007_ := (&models.TaskShape{Name: `Mishap investigation-PBS`}).Stage(stage)
	__TaskShape__00000008_ := (&models.TaskShape{Name: `-WBS`}).Stage(stage)
	__TaskShape__00000009_ := (&models.TaskShape{Name: `Program Investigation Team (PIT) Report-PIT focus`}).Stage(stage)
	__TaskShape__00000010_ := (&models.TaskShape{Name: `-WBS`}).Stage(stage)
	__TaskShape__00000011_ := (&models.TaskShape{Name: `-WBS`}).Stage(stage)

	// insertion point for initialization of values

	__Diagram__00000000_.Name = `Default Diagram`
	__Diagram__00000000_.IsChecked = false
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.ShowPrefix = false
	__Diagram__00000000_.DefaultBoxWidth = 250.000000
	__Diagram__00000000_.DefaultBoxHeigth = 70.000000
	__Diagram__00000000_.IsExpanded = false
	__Diagram__00000000_.ComputedPrefix = ``
	__Diagram__00000000_.IsInRenameMode = false
	__Diagram__00000000_.IsPBSNodeExpanded = true
	__Diagram__00000000_.IsWBSNodeExpanded = true
	__Diagram__00000000_.IsNotesNodeExpanded = true
	__Diagram__00000000_.IsResourcesNodeExpanded = true

	__Diagram__00000001_.Name = `WBS`
	__Diagram__00000001_.IsChecked = false
	__Diagram__00000001_.IsEditable_ = true
	__Diagram__00000001_.ShowPrefix = true
	__Diagram__00000001_.DefaultBoxWidth = 250.000000
	__Diagram__00000001_.DefaultBoxHeigth = 70.000000
	__Diagram__00000001_.IsExpanded = false
	__Diagram__00000001_.ComputedPrefix = ``
	__Diagram__00000001_.IsInRenameMode = false
	__Diagram__00000001_.IsPBSNodeExpanded = false
	__Diagram__00000001_.IsWBSNodeExpanded = true
	__Diagram__00000001_.IsNotesNodeExpanded = false
	__Diagram__00000001_.IsResourcesNodeExpanded = false

	__Diagram__00000002_.Name = `PBS`
	__Diagram__00000002_.IsChecked = false
	__Diagram__00000002_.IsEditable_ = true
	__Diagram__00000002_.ShowPrefix = true
	__Diagram__00000002_.DefaultBoxWidth = 250.000000
	__Diagram__00000002_.DefaultBoxHeigth = 70.000000
	__Diagram__00000002_.IsExpanded = false
	__Diagram__00000002_.ComputedPrefix = ``
	__Diagram__00000002_.IsInRenameMode = false
	__Diagram__00000002_.IsPBSNodeExpanded = true
	__Diagram__00000002_.IsWBSNodeExpanded = false
	__Diagram__00000002_.IsNotesNodeExpanded = true
	__Diagram__00000002_.IsResourcesNodeExpanded = false

	__Diagram__00000003_.Name = `RBS`
	__Diagram__00000003_.IsChecked = false
	__Diagram__00000003_.IsEditable_ = true
	__Diagram__00000003_.ShowPrefix = false
	__Diagram__00000003_.DefaultBoxWidth = 250.000000
	__Diagram__00000003_.DefaultBoxHeigth = 70.000000
	__Diagram__00000003_.IsExpanded = false
	__Diagram__00000003_.ComputedPrefix = ``
	__Diagram__00000003_.IsInRenameMode = false
	__Diagram__00000003_.IsPBSNodeExpanded = true
	__Diagram__00000003_.IsWBSNodeExpanded = false
	__Diagram__00000003_.IsNotesNodeExpanded = false
	__Diagram__00000003_.IsResourcesNodeExpanded = true

	__Diagram__00000004_.Name = `PIT focus`
	__Diagram__00000004_.IsChecked = false
	__Diagram__00000004_.IsEditable_ = true
	__Diagram__00000004_.ShowPrefix = false
	__Diagram__00000004_.DefaultBoxWidth = 250.000000
	__Diagram__00000004_.DefaultBoxHeigth = 70.000000
	__Diagram__00000004_.IsExpanded = false
	__Diagram__00000004_.ComputedPrefix = ``
	__Diagram__00000004_.IsInRenameMode = false
	__Diagram__00000004_.IsPBSNodeExpanded = true
	__Diagram__00000004_.IsWBSNodeExpanded = false
	__Diagram__00000004_.IsNotesNodeExpanded = true
	__Diagram__00000004_.IsResourcesNodeExpanded = false

	__Diagram__00000005_.Name = `PIT Report`
	__Diagram__00000005_.IsChecked = true
	__Diagram__00000005_.IsEditable_ = true
	__Diagram__00000005_.ShowPrefix = false
	__Diagram__00000005_.DefaultBoxWidth = 250.000000
	__Diagram__00000005_.DefaultBoxHeigth = 70.000000
	__Diagram__00000005_.IsExpanded = true
	__Diagram__00000005_.ComputedPrefix = ``
	__Diagram__00000005_.IsInRenameMode = false
	__Diagram__00000005_.IsPBSNodeExpanded = true
	__Diagram__00000005_.IsWBSNodeExpanded = false
	__Diagram__00000005_.IsNotesNodeExpanded = true
	__Diagram__00000005_.IsResourcesNodeExpanded = false

	__Diagram__00000006_.Name = `RCS PBS`
	__Diagram__00000006_.IsChecked = false
	__Diagram__00000006_.IsEditable_ = true
	__Diagram__00000006_.ShowPrefix = true
	__Diagram__00000006_.DefaultBoxWidth = 250.000000
	__Diagram__00000006_.DefaultBoxHeigth = 70.000000
	__Diagram__00000006_.IsExpanded = true
	__Diagram__00000006_.ComputedPrefix = ``
	__Diagram__00000006_.IsInRenameMode = false
	__Diagram__00000006_.IsPBSNodeExpanded = true
	__Diagram__00000006_.IsWBSNodeExpanded = false
	__Diagram__00000006_.IsNotesNodeExpanded = true
	__Diagram__00000006_.IsResourcesNodeExpanded = false

	__Note__00000000_.Name = `CFT ended in march 2025`
	__Note__00000000_.IsExpanded = false
	__Note__00000000_.ComputedPrefix = `1`
	__Note__00000000_.IsInRenameMode = false

	__Note__00000001_.Name = `A thorough review of the STAR report is advised (p14)
`
	__Note__00000001_.IsExpanded = false
	__Note__00000001_.ComputedPrefix = `2`
	__Note__00000001_.IsInRenameMode = false

	__Note__00000002_.Name = `NASA utilized a firm fixed price contracting type for CCtCap. This was a significant shift from the cost-plus contracting for traditional NASA builds of developmental vehicles. These shifts signified that CCP was not only positioned to be an innovative, first-of-its kind program for NASA, but how it interacted with new and traditional space flight industry providers was setup to be significantly distinct and different.  `
	__Note__00000002_.IsExpanded = false
	__Note__00000002_.ComputedPrefix = `3`
	__Note__00000002_.IsInRenameMode = false

	__Note__00000003_.Name = `The CCP 1100 series of requirements were deliberately written at a higher-level, leaving room for provider innovation but there was also room for incorrect/inadequate interpretation by the providers. `
	__Note__00000003_.IsExpanded = false
	__Note__00000003_.ComputedPrefix = `4`
	__Note__00000003_.IsInRenameMode = false

	__Note__00000004_.Name = `The Commercial Provider focused on meeting contractual requirement language resulting in insufficient demonstration across the components/system and ground/flight. `
	__Note__00000004_.IsExpanded = false
	__Note__00000004_.ComputedPrefix = `5`
	__Note__00000004_.IsInRenameMode = false

	__Note__00000005_.Name = `Suppliers’ build quality/variability issues can be hard to exonerate for service modules, which is hardware that is disposed of during re-entry. `
	__Note__00000005_.IsExpanded = false
	__Note__00000005_.ComputedPrefix = `6`
	__Note__00000005_.IsInRenameMode = false

	__Note__00000006_.Name = `The Aerojet Rocketdyne (AR) thermal model included the effects of jet firings, but these effects were not validated by ground testing.

Boeing thermal model did not include the effects of jet firings before CFT. `
	__Note__00000006_.IsExpanded = false
	__Note__00000006_.ComputedPrefix = `7`
	__Note__00000006_.IsInRenameMode = false

	__Note__00000007_.Name = `The thruster performance from OFT1 & OFT2 experienced greater than expected temperatures and continuing to operate lead to a false sense of security of the thruster capability/performance. `
	__Note__00000007_.IsExpanded = false
	__Note__00000007_.ComputedPrefix = `8`
	__Note__00000007_.IsInRenameMode = false

	__Note__00000008_.Name = `Flight instrumentation locations for thermal sensors were limited and in different locations than the locations for RCS Thruster ground firings. `
	__Note__00000008_.IsExpanded = false
	__Note__00000008_.ComputedPrefix = `9`
	__Note__00000008_.IsInRenameMode = false

	__Note__00000009_.Name = `OFT1 & OFT2 investigations did not include RCS/OMAC thruster firings and fault trees were not validated through subsequent ground testing. `
	__Note__00000009_.IsExpanded = false
	__Note__00000009_.ComputedPrefix = `10`
	__Note__00000009_.IsInRenameMode = false

	__Note__00000010_.Name = `For OFT2, NASA/Boeing did not have tools to measure thruster degradation, simply treated the thruster as failed/operational. `
	__Note__00000010_.IsExpanded = false
	__Note__00000010_.ComputedPrefix = `11`
	__Note__00000010_.IsInRenameMode = false

	__Note__00000011_.Name = `Pc (Chamber Pressure).To know if a thruster is actually firing, the spacecraft's computer looks at the pressure sensor inside the combustion chamber (the Pc telemetry). If the pressure shoots up to the expected level, it means combustion is happening and the thruster is pushing. If the pressure stays low or at zero, it assumes the thruster failed.`
	__Note__00000011_.IsExpanded = false
	__Note__00000011_.ComputedPrefix = `12`
	__Note__00000011_.IsInRenameMode = false

	__Note__00000012_.Name = `Pc transducer failed but the engine was able to restart. Boing confidence in the harxware robustness`
	__Note__00000012_.IsExpanded = false
	__Note__00000012_.ComputedPrefix = `13`
	__Note__00000012_.IsInRenameMode = false

	__NoteProductShape__00000000_.Name = `A thorough review of the STAR report is advised to  Starliner Tests and Anomalies Review (STAR) Investigation Report`
	__NoteProductShape__00000000_.StartRatio = 0.500000
	__NoteProductShape__00000000_.EndRatio = 0.500000
	__NoteProductShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000000_.CornerOffsetRatio = 1.680000

	__NoteProductShape__00000001_.Name = `NASA utilized a firm fixed price contracting type for CCtCap. This was a significant shift from the cost-plus contracting for traditional NASA builds of developmental vehicles. These shifts signified that CCP was not only positioned to be an innovative, first-of-its kind program for NASA, but how it interacted with new and traditional space flight industry providers was setup to be significantly distinct and different.   to Commercial Crew Transportation Capability (CCtCap).`
	__NoteProductShape__00000001_.StartRatio = 0.500000
	__NoteProductShape__00000001_.EndRatio = 0.500000
	__NoteProductShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000001_.CornerOffsetRatio = 1.680000

	__NoteProductShape__00000002_.Name = ` to CCP Requirements`
	__NoteProductShape__00000002_.StartRatio = 0.500000
	__NoteProductShape__00000002_.EndRatio = 0.500000
	__NoteProductShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000002_.CornerOffsetRatio = 1.680000

	__NoteProductShape__00000003_.Name = `The Commercial Provider focused on meeting contractual requirement language resulting in insufficient demonstration across the components/system and ground/flight.  to CCP Requirements`
	__NoteProductShape__00000003_.StartRatio = 0.500000
	__NoteProductShape__00000003_.EndRatio = 0.500000
	__NoteProductShape__00000003_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000003_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000003_.CornerOffsetRatio = 1.680000

	__NoteProductShape__00000004_.Name = `Suppliers’ build quality/variability issues can be hard to exonerate for service modules, which is hardware that is disposed of during re-entry.  to CM RCS Thrusters`
	__NoteProductShape__00000004_.StartRatio = 0.500000
	__NoteProductShape__00000004_.EndRatio = 0.500000
	__NoteProductShape__00000004_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000004_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000004_.CornerOffsetRatio = 1.680000

	__NoteProductShape__00000005_.Name = `Suppliers’ build quality/variability issues can be hard to exonerate for service modules, which is hardware that is disposed of during re-entry.  to SM RCS Thrusters`
	__NoteProductShape__00000005_.StartRatio = 0.500000
	__NoteProductShape__00000005_.EndRatio = 0.500000
	__NoteProductShape__00000005_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000005_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000005_.CornerOffsetRatio = 1.680000

	__NoteProductShape__00000006_.Name = `The Aerojet Rocketdyne (AR) thermal model included the effects of jet firings, but these effects were not validated by ground testing.

Boeing thermal model did not include the effects of jet firings before CFT.  to 28 SM RCS Thrusters`
	__NoteProductShape__00000006_.StartRatio = 0.500000
	__NoteProductShape__00000006_.EndRatio = 0.500000
	__NoteProductShape__00000006_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000006_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000006_.CornerOffsetRatio = 1.680000

	__NoteProductShape__00000007_.Name = `The thruster performance from OFT1 & OFT2 experienced greater than expected temperatures and continuing to operate lead to a false sense of security of the thruster capability/performance.  to 28 SM RCS Thrusters`
	__NoteProductShape__00000007_.StartRatio = 0.500000
	__NoteProductShape__00000007_.EndRatio = 0.500000
	__NoteProductShape__00000007_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000007_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000007_.CornerOffsetRatio = 1.680000

	__NoteProductShape__00000008_.Name = `Flight instrumentation locations for thermal sensors were limited and in different locations than the locations for RCS Thruster ground firings.  to Thermal Sensors`
	__NoteProductShape__00000008_.StartRatio = 0.500000
	__NoteProductShape__00000008_.EndRatio = 0.500000
	__NoteProductShape__00000008_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000008_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000008_.CornerOffsetRatio = 1.680000

	__NoteProductShape__00000009_.Name = `Flight instrumentation locations for thermal sensors were limited and in different locations than the locations for RCS Thruster ground firings.  to 28 SM RCS Thrusters`
	__NoteProductShape__00000009_.StartRatio = 0.500000
	__NoteProductShape__00000009_.EndRatio = 0.500000
	__NoteProductShape__00000009_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000009_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000009_.CornerOffsetRatio = 1.680000

	__NoteProductShape__00000010_.Name = `Flight instrumentation locations for thermal sensors were limited and in different locations than the locations for RCS Thruster ground firings.  to 12 CM RCS Thrusters`
	__NoteProductShape__00000010_.StartRatio = 0.500000
	__NoteProductShape__00000010_.EndRatio = 0.500000
	__NoteProductShape__00000010_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000010_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000010_.CornerOffsetRatio = 1.680000

	__NoteProductShape__00000011_.Name = `OFT1 & OFT2 investigations did not include RCS/OMAC thruster firings and fault trees were not validated through subsequent ground testing.  to 28 SM RCS Thrusters (2800 lbf)`
	__NoteProductShape__00000011_.StartRatio = 0.500000
	__NoteProductShape__00000011_.EndRatio = 0.500000
	__NoteProductShape__00000011_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000011_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000011_.CornerOffsetRatio = 1.680000

	__NoteProductShape__00000012_.Name = `For OFT2, NASA/Boeing did not have tools to measure thruster degradation, simply treated the thruster as failed/operational.  to 28 SM RCS Thrusters (2800 lbf)`
	__NoteProductShape__00000012_.StartRatio = 0.500000
	__NoteProductShape__00000012_.EndRatio = 0.500000
	__NoteProductShape__00000012_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000012_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000012_.CornerOffsetRatio = 1.680000

	__NoteProductShape__00000013_.Name = `Pc (Chamber Pressure) to 12 CM RCS Thrusters (1200 lbf)`
	__NoteProductShape__00000013_.StartRatio = 0.500000
	__NoteProductShape__00000013_.EndRatio = 0.500000
	__NoteProductShape__00000013_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000013_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000013_.CornerOffsetRatio = 1.680000

	__NoteProductShape__00000014_.Name = `Pc (Chamber Pressure).To know if a thruster is actually firing, the spacecraft's computer looks at the pressure sensor inside the combustion chamber (the Pc telemetry). If the pressure shoots up to the expected level, it means combustion is happening and the thruster is pushing. If the pressure stays low or at zero, it assumes the thruster failed. to Transducer / pressure sensor`
	__NoteProductShape__00000014_.StartRatio = 0.500000
	__NoteProductShape__00000014_.EndRatio = 0.500000
	__NoteProductShape__00000014_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000014_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000014_.CornerOffsetRatio = 1.680000

	__NoteProductShape__00000015_.Name = `Pc transducer failed but the engine was able to restart. Boing confidence in the harxware robustness to Transducer / pressure sensor (Pc)`
	__NoteProductShape__00000015_.StartRatio = 0.500000
	__NoteProductShape__00000015_.EndRatio = 0.500000
	__NoteProductShape__00000015_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000015_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000015_.CornerOffsetRatio = 1.680000

	__NoteShape__00000000_.Name = `-Default Diagram`
	__NoteShape__00000000_.IsExpanded = false
	__NoteShape__00000000_.X = 92.441418
	__NoteShape__00000000_.Y = 1047.751617
	__NoteShape__00000000_.Width = 250.000000
	__NoteShape__00000000_.Height = 70.000000

	__NoteShape__00000001_.Name = `-PIT focus`
	__NoteShape__00000001_.IsExpanded = false
	__NoteShape__00000001_.X = 892.606649
	__NoteShape__00000001_.Y = 74.204707
	__NoteShape__00000001_.Width = 250.000000
	__NoteShape__00000001_.Height = 70.000000

	__NoteShape__00000002_.Name = `-PBS`
	__NoteShape__00000002_.IsExpanded = false
	__NoteShape__00000002_.X = 465.290428
	__NoteShape__00000002_.Y = 281.921876
	__NoteShape__00000002_.Width = 294.000000
	__NoteShape__00000002_.Height = 266.999985

	__NoteShape__00000003_.Name = `-PIT Report`
	__NoteShape__00000003_.IsExpanded = false
	__NoteShape__00000003_.X = 1285.673362
	__NoteShape__00000003_.Y = 221.503075
	__NoteShape__00000003_.Width = 250.000000
	__NoteShape__00000003_.Height = 177.000000

	__NoteShape__00000004_.Name = `-PIT Report`
	__NoteShape__00000004_.IsExpanded = false
	__NoteShape__00000004_.X = 852.720613
	__NoteShape__00000004_.Y = 286.622045
	__NoteShape__00000004_.Width = 250.000000
	__NoteShape__00000004_.Height = 143.999939

	__NoteShape__00000005_.Name = `-RCS PBS`
	__NoteShape__00000005_.IsExpanded = false
	__NoteShape__00000005_.X = 512.089724
	__NoteShape__00000005_.Y = 893.092030
	__NoteShape__00000005_.Width = 250.000000
	__NoteShape__00000005_.Height = 194.999939

	__NoteShape__00000006_.Name = `-RCS PBS`
	__NoteShape__00000006_.IsExpanded = false
	__NoteShape__00000006_.X = 209.153628
	__NoteShape__00000006_.Y = 729.461448
	__NoteShape__00000006_.Width = 250.000000
	__NoteShape__00000006_.Height = 196.000000

	__NoteShape__00000007_.Name = `-RCS PBS`
	__NoteShape__00000007_.IsExpanded = false
	__NoteShape__00000007_.X = 596.510371
	__NoteShape__00000007_.Y = 678.471309
	__NoteShape__00000007_.Width = 269.000000
	__NoteShape__00000007_.Height = 181.999939

	__NoteShape__00000008_.Name = `-RCS PBS`
	__NoteShape__00000008_.IsExpanded = false
	__NoteShape__00000008_.X = 943.375080
	__NoteShape__00000008_.Y = 524.315922
	__NoteShape__00000008_.Width = 250.000000
	__NoteShape__00000008_.Height = 144.000000

	__NoteShape__00000009_.Name = `-RCS PBS`
	__NoteShape__00000009_.IsExpanded = false
	__NoteShape__00000009_.X = 880.446589
	__NoteShape__00000009_.Y = 755.302818
	__NoteShape__00000009_.Width = 250.000000
	__NoteShape__00000009_.Height = 154.000061

	__NoteShape__00000010_.Name = `-RCS PBS`
	__NoteShape__00000010_.IsExpanded = false
	__NoteShape__00000010_.X = 868.189303
	__NoteShape__00000010_.Y = 1020.200971
	__NoteShape__00000010_.Width = 250.000000
	__NoteShape__00000010_.Height = 114.000000

	__NoteShape__00000011_.Name = `-RCS PBS`
	__NoteShape__00000011_.IsExpanded = false
	__NoteShape__00000011_.X = 45.256122
	__NoteShape__00000011_.Y = 950.529205
	__NoteShape__00000011_.Width = 250.000000
	__NoteShape__00000011_.Height = 227.000061

	__NoteShape__00000012_.Name = `-RCS PBS`
	__NoteShape__00000012_.IsExpanded = false
	__NoteShape__00000012_.X = 328.566920
	__NoteShape__00000012_.Y = 1060.052255
	__NoteShape__00000012_.Width = 250.000000
	__NoteShape__00000012_.Height = 121.000000

	__NoteTaskShape__00000000_.Name = `CFT ended in march 2025 to Starliner Crewed Flight Test (CFT)`
	__NoteTaskShape__00000000_.StartRatio = 0.500000
	__NoteTaskShape__00000000_.EndRatio = 0.500000
	__NoteTaskShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteTaskShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteTaskShape__00000000_.CornerOffsetRatio = 1.680000

	__Product__00000001_.Name = `Dragon`
	__Product__00000001_.Description = ``
	__Product__00000001_.IsExpanded = false
	__Product__00000001_.ComputedPrefix = `1.1.2`
	__Product__00000001_.IsInRenameMode = false
	__Product__00000001_.IsProducersNodeExpanded = false
	__Product__00000001_.IsConsumersNodeExpanded = false

	__Product__00000002_.Name = `CST-100 Starliner`
	__Product__00000002_.Description = ``
	__Product__00000002_.IsExpanded = false
	__Product__00000002_.ComputedPrefix = `1.1.1`
	__Product__00000002_.IsInRenameMode = false
	__Product__00000002_.IsProducersNodeExpanded = false
	__Product__00000002_.IsConsumersNodeExpanded = false

	__Product__00000003_.Name = `Reports`
	__Product__00000003_.Description = ``
	__Product__00000003_.IsExpanded = false
	__Product__00000003_.ComputedPrefix = `2`
	__Product__00000003_.IsInRenameMode = false
	__Product__00000003_.IsProducersNodeExpanded = false
	__Product__00000003_.IsConsumersNodeExpanded = false

	__Product__00000004_.Name = ` Starliner Tests and Anomalies Review (STAR) Investigation Report`
	__Product__00000004_.Description = ``
	__Product__00000004_.IsExpanded = false
	__Product__00000004_.ComputedPrefix = `2.1`
	__Product__00000004_.IsInRenameMode = false
	__Product__00000004_.IsProducersNodeExpanded = false
	__Product__00000004_.IsConsumersNodeExpanded = false

	__Product__00000005_.Name = `Program Investigation Team (PIT) Report`
	__Product__00000005_.Description = ``
	__Product__00000005_.IsExpanded = false
	__Product__00000005_.ComputedPrefix = `2.2`
	__Product__00000005_.IsInRenameMode = false
	__Product__00000005_.IsProducersNodeExpanded = false
	__Product__00000005_.IsConsumersNodeExpanded = false

	__Product__00000006_.Name = `Commercial Crew Transportation Capability (CCtCap).`
	__Product__00000006_.Description = ``
	__Product__00000006_.IsExpanded = false
	__Product__00000006_.ComputedPrefix = `1.1`
	__Product__00000006_.IsInRenameMode = false
	__Product__00000006_.IsProducersNodeExpanded = false
	__Product__00000006_.IsConsumersNodeExpanded = false

	__Product__00000007_.Name = `NASA Assets/Capabities`
	__Product__00000007_.Description = ``
	__Product__00000007_.IsExpanded = false
	__Product__00000007_.ComputedPrefix = `1`
	__Product__00000007_.IsInRenameMode = false
	__Product__00000007_.IsProducersNodeExpanded = false
	__Product__00000007_.IsConsumersNodeExpanded = false

	__Product__00000008_.Name = `ISS`
	__Product__00000008_.Description = ``
	__Product__00000008_.IsExpanded = false
	__Product__00000008_.ComputedPrefix = `1.2`
	__Product__00000008_.IsInRenameMode = false
	__Product__00000008_.IsProducersNodeExpanded = false
	__Product__00000008_.IsConsumersNodeExpanded = false

	__Product__00000009_.Name = `3 Commercial Crew Program (CCP) Background`
	__Product__00000009_.Description = ``
	__Product__00000009_.IsExpanded = false
	__Product__00000009_.ComputedPrefix = `2.2.2`
	__Product__00000009_.IsInRenameMode = false
	__Product__00000009_.IsProducersNodeExpanded = false
	__Product__00000009_.IsConsumersNodeExpanded = false

	__Product__00000010_.Name = `CCP Requirements`
	__Product__00000010_.Description = ``
	__Product__00000010_.IsExpanded = false
	__Product__00000010_.ComputedPrefix = `1.1.3`
	__Product__00000010_.IsInRenameMode = false
	__Product__00000010_.IsProducersNodeExpanded = false
	__Product__00000010_.IsConsumersNodeExpanded = false

	__Product__00000011_.Name = `Crew Module (CM)`
	__Product__00000011_.Description = ``
	__Product__00000011_.IsExpanded = false
	__Product__00000011_.ComputedPrefix = `1.1.1.1`
	__Product__00000011_.IsInRenameMode = false
	__Product__00000011_.IsProducersNodeExpanded = false
	__Product__00000011_.IsConsumersNodeExpanded = false

	__Product__00000012_.Name = `CM Reaction Control System (RCS)`
	__Product__00000012_.Description = ``
	__Product__00000012_.IsExpanded = false
	__Product__00000012_.ComputedPrefix = `1.1.1.1.1`
	__Product__00000012_.IsInRenameMode = false
	__Product__00000012_.IsProducersNodeExpanded = false
	__Product__00000012_.IsConsumersNodeExpanded = false

	__Product__00000013_.Name = `12 CM RCS Thrusters (1200 lbf)`
	__Product__00000013_.Description = ``
	__Product__00000013_.IsExpanded = false
	__Product__00000013_.ComputedPrefix = `1.1.1.1.1.1`
	__Product__00000013_.IsInRenameMode = false
	__Product__00000013_.IsProducersNodeExpanded = false
	__Product__00000013_.IsConsumersNodeExpanded = false

	__Product__00000014_.Name = `Service Module (SM)`
	__Product__00000014_.Description = ``
	__Product__00000014_.IsExpanded = false
	__Product__00000014_.ComputedPrefix = `1.1.1.2`
	__Product__00000014_.IsInRenameMode = false
	__Product__00000014_.IsProducersNodeExpanded = false
	__Product__00000014_.IsConsumersNodeExpanded = false

	__Product__00000015_.Name = `SM Reaction Control System (SM RCS)`
	__Product__00000015_.Description = ``
	__Product__00000015_.IsExpanded = false
	__Product__00000015_.ComputedPrefix = `1.1.1.2.1`
	__Product__00000015_.IsInRenameMode = false
	__Product__00000015_.IsProducersNodeExpanded = false
	__Product__00000015_.IsConsumersNodeExpanded = false

	__Product__00000016_.Name = `28 SM RCS Thrusters (2800 lbf)`
	__Product__00000016_.Description = ``
	__Product__00000016_.IsExpanded = false
	__Product__00000016_.ComputedPrefix = `1.1.1.2.1.1`
	__Product__00000016_.IsInRenameMode = false
	__Product__00000016_.IsProducersNodeExpanded = false
	__Product__00000016_.IsConsumersNodeExpanded = false

	__Product__00000017_.Name = `2 Additional Investigations `
	__Product__00000017_.Description = ``
	__Product__00000017_.IsExpanded = false
	__Product__00000017_.ComputedPrefix = `2.2.1`
	__Product__00000017_.IsInRenameMode = false
	__Product__00000017_.IsProducersNodeExpanded = false
	__Product__00000017_.IsConsumersNodeExpanded = false

	__Product__00000018_.Name = `STAR Report Summary and Findings `
	__Product__00000018_.Description = ``
	__Product__00000018_.IsExpanded = false
	__Product__00000018_.ComputedPrefix = `2.2.1.1`
	__Product__00000018_.IsInRenameMode = false
	__Product__00000018_.IsProducersNodeExpanded = false
	__Product__00000018_.IsConsumersNodeExpanded = false

	__Product__00000019_.Name = `Thermal Sensors`
	__Product__00000019_.Description = ``
	__Product__00000019_.IsExpanded = false
	__Product__00000019_.ComputedPrefix = `1.1.1.3`
	__Product__00000019_.IsInRenameMode = false
	__Product__00000019_.IsProducersNodeExpanded = false
	__Product__00000019_.IsConsumersNodeExpanded = false

	__Product__00000020_.Name = `Dog House`
	__Product__00000020_.Description = ``
	__Product__00000020_.IsExpanded = false
	__Product__00000020_.ComputedPrefix = `1.1.1.1.1.1.1`
	__Product__00000020_.IsInRenameMode = false
	__Product__00000020_.IsProducersNodeExpanded = false
	__Product__00000020_.IsConsumersNodeExpanded = false

	__Product__00000021_.Name = `Orbital Maneuvering and Attitude Control (OMAC), 30 000 lbf`
	__Product__00000021_.Description = ``
	__Product__00000021_.IsExpanded = false
	__Product__00000021_.ComputedPrefix = `1.1.1.4`
	__Product__00000021_.IsInRenameMode = false
	__Product__00000021_.IsProducersNodeExpanded = false
	__Product__00000021_.IsConsumersNodeExpanded = false

	__Product__00000022_.Name = `Boeing’s Enterprise Root Cause/Corrective Actions (eRCCA)`
	__Product__00000022_.Description = ``
	__Product__00000022_.IsExpanded = false
	__Product__00000022_.ComputedPrefix = `2.3`
	__Product__00000022_.IsInRenameMode = false
	__Product__00000022_.IsProducersNodeExpanded = false
	__Product__00000022_.IsConsumersNodeExpanded = false

	__Product__00000023_.Name = `3.1   Orbital Flight Test (OFT) Summary`
	__Product__00000023_.Description = ``
	__Product__00000023_.IsExpanded = false
	__Product__00000023_.ComputedPrefix = `2.2.2.1`
	__Product__00000023_.IsInRenameMode = false
	__Product__00000023_.IsProducersNodeExpanded = false
	__Product__00000023_.IsConsumersNodeExpanded = false

	__Product__00000024_.Name = `3.2   Orbital Flight Test 2 (OFT-2) Summary`
	__Product__00000024_.Description = ``
	__Product__00000024_.IsExpanded = false
	__Product__00000024_.ComputedPrefix = `2.2.2.2`
	__Product__00000024_.IsInRenameMode = false
	__Product__00000024_.IsProducersNodeExpanded = false
	__Product__00000024_.IsConsumersNodeExpanded = false

	__Product__00000025_.Name = `Flight Software (FSW)`
	__Product__00000025_.Description = ``
	__Product__00000025_.IsExpanded = false
	__Product__00000025_.ComputedPrefix = `1.1.1.5`
	__Product__00000025_.IsInRenameMode = false
	__Product__00000025_.IsProducersNodeExpanded = false
	__Product__00000025_.IsConsumersNodeExpanded = false

	__Product__00000026_.Name = `Fault Detection, Isolation, and Recovery (FDIR)`
	__Product__00000026_.Description = ``
	__Product__00000026_.IsExpanded = false
	__Product__00000026_.ComputedPrefix = `1.1.1.6`
	__Product__00000026_.IsInRenameMode = false
	__Product__00000026_.IsProducersNodeExpanded = false
	__Product__00000026_.IsConsumersNodeExpanded = false

	__Product__00000027_.Name = `3.3 Comparing SM RCS Thrusters Triggering Fail-Off FDIR on OFT1/OFT2`
	__Product__00000027_.Description = ``
	__Product__00000027_.IsExpanded = false
	__Product__00000027_.ComputedPrefix = `2.2.2.3`
	__Product__00000027_.IsInRenameMode = false
	__Product__00000027_.IsProducersNodeExpanded = false
	__Product__00000027_.IsConsumersNodeExpanded = false

	__Product__00000028_.Name = `Transducer / pressure sensor (Pc)`
	__Product__00000028_.Description = ``
	__Product__00000028_.IsExpanded = false
	__Product__00000028_.ComputedPrefix = `1.1.1.1.1.1.2`
	__Product__00000028_.IsInRenameMode = false
	__Product__00000028_.IsProducersNodeExpanded = false
	__Product__00000028_.IsConsumersNodeExpanded = false

	__Product__00000029_.Name = `4     Technical Root Cause Analysis (RCA) and Findings`
	__Product__00000029_.Description = ``
	__Product__00000029_.IsExpanded = false
	__Product__00000029_.ComputedPrefix = `2.2.3`
	__Product__00000029_.IsInRenameMode = false
	__Product__00000029_.IsProducersNodeExpanded = false
	__Product__00000029_.IsConsumersNodeExpanded = false

	__ProductCompositionShape__00000000_.Name = `Reports to `
	__ProductCompositionShape__00000000_.StartRatio = 0.500000
	__ProductCompositionShape__00000000_.EndRatio = 0.500000
	__ProductCompositionShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000000_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000001_.Name = `Reports to `
	__ProductCompositionShape__00000001_.StartRatio = 0.500000
	__ProductCompositionShape__00000001_.EndRatio = 0.500000
	__ProductCompositionShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000001_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000002_.Name = `Commercial Crew Transportation Capability (CCtCap). to Starliner`
	__ProductCompositionShape__00000002_.StartRatio = 0.500000
	__ProductCompositionShape__00000002_.EndRatio = 0.500000
	__ProductCompositionShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000002_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000003_.Name = `Commercial Crew Transportation Capability (CCtCap). to Dragon`
	__ProductCompositionShape__00000003_.StartRatio = 0.500000
	__ProductCompositionShape__00000003_.EndRatio = 0.500000
	__ProductCompositionShape__00000003_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000003_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000003_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000004_.Name = `NASA Assets/Capabities to Commercial Crew Transportation Capability (CCtCap).`
	__ProductCompositionShape__00000004_.StartRatio = 0.500000
	__ProductCompositionShape__00000004_.EndRatio = 0.500000
	__ProductCompositionShape__00000004_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000004_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000004_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000005_.Name = `Commercial Crew Transportation Capability (CCtCap). to `
	__ProductCompositionShape__00000005_.StartRatio = 0.501832
	__ProductCompositionShape__00000005_.EndRatio = 0.300650
	__ProductCompositionShape__00000005_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000005_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000005_.CornerOffsetRatio = -0.228174

	__ProductCompositionShape__00000006_.Name = `Program Investigation Team (PIT) Report to `
	__ProductCompositionShape__00000006_.StartRatio = 0.500000
	__ProductCompositionShape__00000006_.EndRatio = 0.460874
	__ProductCompositionShape__00000006_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000006_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000006_.CornerOffsetRatio = 0.192000

	__ProductCompositionShape__00000007_.Name = `CM to `
	__ProductCompositionShape__00000007_.StartRatio = 0.500000
	__ProductCompositionShape__00000007_.EndRatio = 0.500000
	__ProductCompositionShape__00000007_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000007_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000007_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000008_.Name = `RCS to `
	__ProductCompositionShape__00000008_.StartRatio = 0.500000
	__ProductCompositionShape__00000008_.EndRatio = 0.500000
	__ProductCompositionShape__00000008_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000008_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000008_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000009_.Name = `CST-100 Starliner to CM`
	__ProductCompositionShape__00000009_.StartRatio = 0.500000
	__ProductCompositionShape__00000009_.EndRatio = 0.500000
	__ProductCompositionShape__00000009_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000009_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000009_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000010_.Name = `CST-100 Starliner to `
	__ProductCompositionShape__00000010_.StartRatio = 0.500000
	__ProductCompositionShape__00000010_.EndRatio = 0.500000
	__ProductCompositionShape__00000010_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000010_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000010_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000011_.Name = `Service Module (SM) to `
	__ProductCompositionShape__00000011_.StartRatio = 0.500000
	__ProductCompositionShape__00000011_.EndRatio = 0.500000
	__ProductCompositionShape__00000011_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000011_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000011_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000012_.Name = `SM Reaction Control System (SM RCS) to `
	__ProductCompositionShape__00000012_.StartRatio = 0.500000
	__ProductCompositionShape__00000012_.EndRatio = 0.500000
	__ProductCompositionShape__00000012_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000012_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000012_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000013_.Name = `Program Investigation Team (PIT) Report to `
	__ProductCompositionShape__00000013_.StartRatio = 0.500000
	__ProductCompositionShape__00000013_.EndRatio = 0.498374
	__ProductCompositionShape__00000013_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000013_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000013_.CornerOffsetRatio = 0.264000

	__ProductCompositionShape__00000014_.Name = `Chap 2 Additional Investigations  to `
	__ProductCompositionShape__00000014_.StartRatio = 0.500000
	__ProductCompositionShape__00000014_.EndRatio = 0.500000
	__ProductCompositionShape__00000014_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000014_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000014_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000015_.Name = `Crew Module (CM) to `
	__ProductCompositionShape__00000015_.StartRatio = 0.549056
	__ProductCompositionShape__00000015_.EndRatio = 0.345913
	__ProductCompositionShape__00000015_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000015_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000015_.CornerOffsetRatio = 5.732012

	__ProductCompositionShape__00000016_.Name = `12 CM RCS Thrusters to `
	__ProductCompositionShape__00000016_.StartRatio = 0.500000
	__ProductCompositionShape__00000016_.EndRatio = 0.500000
	__ProductCompositionShape__00000016_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000016_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000016_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000017_.Name = `CST-100 Starliner to `
	__ProductCompositionShape__00000017_.StartRatio = 0.577628
	__ProductCompositionShape__00000017_.EndRatio = 0.306199
	__ProductCompositionShape__00000017_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000017_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000017_.CornerOffsetRatio = 5.744012

	__ProductCompositionShape__00000018_.Name = `Reports to Boeing’s Enterprise Root Cause/Corrective Actions (eRCCA)`
	__ProductCompositionShape__00000018_.StartRatio = 0.500000
	__ProductCompositionShape__00000018_.EndRatio = 0.000000
	__ProductCompositionShape__00000018_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000018_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000018_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000019_.Name = `Chap3 Commercial Crew Program (CCP) Background to `
	__ProductCompositionShape__00000019_.StartRatio = 0.500000
	__ProductCompositionShape__00000019_.EndRatio = 0.500000
	__ProductCompositionShape__00000019_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000019_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000019_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000020_.Name = `Chap3 Commercial Crew Program (CCP) Background to `
	__ProductCompositionShape__00000020_.StartRatio = 0.500000
	__ProductCompositionShape__00000020_.EndRatio = 0.500000
	__ProductCompositionShape__00000020_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000020_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000020_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000021_.Name = `CST-100 Starliner to `
	__ProductCompositionShape__00000021_.StartRatio = 0.520485
	__ProductCompositionShape__00000021_.EndRatio = 0.363342
	__ProductCompositionShape__00000021_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000021_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000021_.CornerOffsetRatio = 5.740012

	__ProductCompositionShape__00000022_.Name = `CST-100 Starliner to `
	__ProductCompositionShape__00000022_.StartRatio = 0.506199
	__ProductCompositionShape__00000022_.EndRatio = 0.149056
	__ProductCompositionShape__00000022_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000022_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000022_.CornerOffsetRatio = 5.740012

	__ProductCompositionShape__00000023_.Name = `Chap3 Commercial Crew Program (CCP) Background to `
	__ProductCompositionShape__00000023_.StartRatio = 0.500000
	__ProductCompositionShape__00000023_.EndRatio = 0.500000
	__ProductCompositionShape__00000023_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000023_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000023_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000024_.Name = `12 CM RCS Thrusters (1200 lbf) to `
	__ProductCompositionShape__00000024_.StartRatio = 0.500000
	__ProductCompositionShape__00000024_.EndRatio = 0.500000
	__ProductCompositionShape__00000024_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000024_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000024_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000025_.Name = `Chap3 Commercial Crew Program (CCP) Background to `
	__ProductCompositionShape__00000025_.StartRatio = 0.500000
	__ProductCompositionShape__00000025_.EndRatio = 0.526945
	__ProductCompositionShape__00000025_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000025_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000025_.CornerOffsetRatio = 0.212000

	__ProductShape__00000002_.Name = `Dragon-PBS`
	__ProductShape__00000002_.IsExpanded = false
	__ProductShape__00000002_.X = 248.020302
	__ProductShape__00000002_.Y = 747.039767
	__ProductShape__00000002_.Width = 250.000000
	__ProductShape__00000002_.Height = 70.000000

	__ProductShape__00000003_.Name = `Starliner-PBS`
	__ProductShape__00000003_.IsExpanded = false
	__ProductShape__00000003_.X = 73.852191
	__ProductShape__00000003_.Y = 583.054452
	__ProductShape__00000003_.Width = 250.000000
	__ProductShape__00000003_.Height = 70.000000

	__ProductShape__00000004_.Name = `-PBS`
	__ProductShape__00000004_.IsExpanded = false
	__ProductShape__00000004_.X = 841.980809
	__ProductShape__00000004_.Y = 105.166614
	__ProductShape__00000004_.Width = 250.000000
	__ProductShape__00000004_.Height = 70.000000

	__ProductShape__00000005_.Name = `-PBS`
	__ProductShape__00000005_.IsExpanded = false
	__ProductShape__00000005_.X = 841.980809
	__ProductShape__00000005_.Y = 245.166614
	__ProductShape__00000005_.Width = 250.000000
	__ProductShape__00000005_.Height = 70.000000

	__ProductShape__00000006_.Name = `-PBS`
	__ProductShape__00000006_.IsExpanded = false
	__ProductShape__00000006_.X = 1141.980809
	__ProductShape__00000006_.Y = 245.166614
	__ProductShape__00000006_.Width = 250.000000
	__ProductShape__00000006_.Height = 70.000000

	__ProductShape__00000007_.Name = `Program Investigation Team (PIT) Report-PIT focus`
	__ProductShape__00000007_.IsExpanded = false
	__ProductShape__00000007_.X = 912.017987
	__ProductShape__00000007_.Y = 275.867255
	__ProductShape__00000007_.Width = 250.000000
	__ProductShape__00000007_.Height = 70.000000

	__ProductShape__00000008_.Name = ` Starliner Tests and Anomalies Review (STAR) Investigation Report-PIT focus`
	__ProductShape__00000008_.IsExpanded = false
	__ProductShape__00000008_.X = 499.019526
	__ProductShape__00000008_.Y = 77.170905
	__ProductShape__00000008_.Width = 250.000000
	__ProductShape__00000008_.Height = 70.000000

	__ProductShape__00000009_.Name = `-PBS`
	__ProductShape__00000009_.IsExpanded = false
	__ProductShape__00000009_.X = 68.071765
	__ProductShape__00000009_.Y = 367.961411
	__ProductShape__00000009_.Width = 250.000000
	__ProductShape__00000009_.Height = 70.000000

	__ProductShape__00000010_.Name = `-PBS`
	__ProductShape__00000010_.IsExpanded = false
	__ProductShape__00000010_.X = 101.619741
	__ProductShape__00000010_.Y = 185.878693
	__ProductShape__00000010_.Width = 250.000000
	__ProductShape__00000010_.Height = 70.000000

	__ProductShape__00000011_.Name = `-PBS`
	__ProductShape__00000011_.IsExpanded = false
	__ProductShape__00000011_.X = 129.071765
	__ProductShape__00000011_.Y = 982.961381
	__ProductShape__00000011_.Width = 250.000000
	__ProductShape__00000011_.Height = 70.000000

	__ProductShape__00000012_.Name = `Program Investigation Team (PIT) Report-PIT Report`
	__ProductShape__00000012_.IsExpanded = false
	__ProductShape__00000012_.X = 41.516205
	__ProductShape__00000012_.Y = 128.119696
	__ProductShape__00000012_.Width = 250.000000
	__ProductShape__00000012_.Height = 70.000000

	__ProductShape__00000013_.Name = `-PIT Report`
	__ProductShape__00000013_.IsExpanded = false
	__ProductShape__00000013_.X = 190.516205
	__ProductShape__00000013_.Y = 540.119696
	__ProductShape__00000013_.Width = 276.000000
	__ProductShape__00000013_.Height = 70.000000

	__ProductShape__00000014_.Name = `-PIT Report`
	__ProductShape__00000014_.IsExpanded = false
	__ProductShape__00000014_.X = 766.767652
	__ProductShape__00000014_.Y = 110.422064
	__ProductShape__00000014_.Width = 250.000000
	__ProductShape__00000014_.Height = 70.000000

	__ProductShape__00000015_.Name = `-RCS PBS`
	__ProductShape__00000015_.IsExpanded = false
	__ProductShape__00000015_.X = 34.373886
	__ProductShape__00000015_.Y = 167.792999
	__ProductShape__00000015_.Width = 250.000000
	__ProductShape__00000015_.Height = 70.000000

	__ProductShape__00000016_.Name = `-RCS PBS`
	__ProductShape__00000016_.IsExpanded = false
	__ProductShape__00000016_.X = 34.373886
	__ProductShape__00000016_.Y = 307.792999
	__ProductShape__00000016_.Width = 250.000000
	__ProductShape__00000016_.Height = 70.000000

	__ProductShape__00000017_.Name = `-RCS PBS`
	__ProductShape__00000017_.IsExpanded = false
	__ProductShape__00000017_.X = 34.373886
	__ProductShape__00000017_.Y = 447.792999
	__ProductShape__00000017_.Width = 250.000000
	__ProductShape__00000017_.Height = 70.000000

	__ProductShape__00000018_.Name = `CST-100 Starliner-RCS PBS`
	__ProductShape__00000018_.IsExpanded = false
	__ProductShape__00000018_.X = 33.736439
	__ProductShape__00000018_.Y = 23.572991
	__ProductShape__00000018_.Width = 250.000000
	__ProductShape__00000018_.Height = 70.000000

	__ProductShape__00000019_.Name = `-RCS PBS`
	__ProductShape__00000019_.IsExpanded = false
	__ProductShape__00000019_.X = 595.736439
	__ProductShape__00000019_.Y = 182.573022
	__ProductShape__00000019_.Width = 250.000000
	__ProductShape__00000019_.Height = 70.000000

	__ProductShape__00000020_.Name = `-RCS PBS`
	__ProductShape__00000020_.IsExpanded = false
	__ProductShape__00000020_.X = 595.736439
	__ProductShape__00000020_.Y = 322.573022
	__ProductShape__00000020_.Width = 250.000000
	__ProductShape__00000020_.Height = 70.000000

	__ProductShape__00000021_.Name = `-RCS PBS`
	__ProductShape__00000021_.IsExpanded = false
	__ProductShape__00000021_.X = 595.736439
	__ProductShape__00000021_.Y = 462.573022
	__ProductShape__00000021_.Width = 250.000000
	__ProductShape__00000021_.Height = 70.000000

	__ProductShape__00000022_.Name = `-PIT Report`
	__ProductShape__00000022_.IsExpanded = false
	__ProductShape__00000022_.X = 207.516205
	__ProductShape__00000022_.Y = 296.119696
	__ProductShape__00000022_.Width = 250.000000
	__ProductShape__00000022_.Height = 70.000000

	__ProductShape__00000023_.Name = `-PIT Report`
	__ProductShape__00000023_.IsExpanded = false
	__ProductShape__00000023_.X = 340.516205
	__ProductShape__00000023_.Y = 440.119696
	__ProductShape__00000023_.Width = 285.000000
	__ProductShape__00000023_.Height = 70.000000

	__ProductShape__00000024_.Name = `-RCS PBS`
	__ProductShape__00000024_.IsExpanded = false
	__ProductShape__00000024_.X = 949.373825
	__ProductShape__00000024_.Y = 158.793044
	__ProductShape__00000024_.Width = 250.000000
	__ProductShape__00000024_.Height = 70.000000

	__ProductShape__00000025_.Name = `-RCS PBS`
	__ProductShape__00000025_.IsExpanded = false
	__ProductShape__00000025_.X = 34.373886
	__ProductShape__00000025_.Y = 587.792999
	__ProductShape__00000025_.Width = 250.000000
	__ProductShape__00000025_.Height = 70.000000

	__ProductShape__00000026_.Name = `-RCS PBS`
	__ProductShape__00000026_.IsExpanded = false
	__ProductShape__00000026_.X = 1173.736439
	__ProductShape__00000026_.Y = 250.572976
	__ProductShape__00000026_.Width = 250.000000
	__ProductShape__00000026_.Height = 70.000000

	__ProductShape__00000027_.Name = `Boeing’s Enterprise Root Cause/Corrective Actions (eRCCA)-PBS`
	__ProductShape__00000027_.IsExpanded = false
	__ProductShape__00000027_.X = 827.950526
	__ProductShape__00000027_.Y = 520.551130
	__ProductShape__00000027_.Width = 250.000000
	__ProductShape__00000027_.Height = 86.000000

	__ProductShape__00000028_.Name = `-PIT Report`
	__ProductShape__00000028_.IsExpanded = false
	__ProductShape__00000028_.X = 331.516205
	__ProductShape__00000028_.Y = 687.119696
	__ProductShape__00000028_.Width = 250.000000
	__ProductShape__00000028_.Height = 70.000000

	__ProductShape__00000029_.Name = `-PIT Report`
	__ProductShape__00000029_.IsExpanded = false
	__ProductShape__00000029_.X = 631.516205
	__ProductShape__00000029_.Y = 687.119696
	__ProductShape__00000029_.Width = 250.000000
	__ProductShape__00000029_.Height = 70.000000

	__ProductShape__00000030_.Name = `-RCS PBS`
	__ProductShape__00000030_.IsExpanded = false
	__ProductShape__00000030_.X = 1171.736439
	__ProductShape__00000030_.Y = 335.572976
	__ProductShape__00000030_.Width = 250.000000
	__ProductShape__00000030_.Height = 70.000000

	__ProductShape__00000031_.Name = `-RCS PBS`
	__ProductShape__00000031_.IsExpanded = false
	__ProductShape__00000031_.X = 1174.736439
	__ProductShape__00000031_.Y = 421.572976
	__ProductShape__00000031_.Width = 250.000000
	__ProductShape__00000031_.Height = 70.000000

	__ProductShape__00000032_.Name = `-PIT Report`
	__ProductShape__00000032_.IsExpanded = false
	__ProductShape__00000032_.X = 931.516205
	__ProductShape__00000032_.Y = 687.119696
	__ProductShape__00000032_.Width = 250.000000
	__ProductShape__00000032_.Height = 70.000000

	__ProductShape__00000033_.Name = `-RCS PBS`
	__ProductShape__00000033_.IsExpanded = false
	__ProductShape__00000033_.X = 334.373886
	__ProductShape__00000033_.Y = 587.792999
	__ProductShape__00000033_.Width = 250.000000
	__ProductShape__00000033_.Height = 70.000000

	__ProductShape__00000034_.Name = `-PIT Report`
	__ProductShape__00000034_.IsExpanded = false
	__ProductShape__00000034_.X = 194.516205
	__ProductShape__00000034_.Y = 800.119696
	__ProductShape__00000034_.Width = 250.000000
	__ProductShape__00000034_.Height = 70.000000

	__Project__00000000_.Name = `Startliner Mishape Report`
	__Project__00000000_.IsExpanded = true
	__Project__00000000_.ComputedPrefix = ``
	__Project__00000000_.IsInRenameMode = false

	__Resource__00000000_.Name = `Program Investigation Team (PIT)`
	__Resource__00000000_.Description = ``
	__Resource__00000000_.IsExpanded = false
	__Resource__00000000_.ComputedPrefix = `1.2`
	__Resource__00000000_.IsInRenameMode = false

	__Resource__00000001_.Name = `Barry "Butch" Wilmore`
	__Resource__00000001_.Description = ``
	__Resource__00000001_.IsExpanded = false
	__Resource__00000001_.ComputedPrefix = `1.1.1.1`
	__Resource__00000001_.IsInRenameMode = false

	__Resource__00000002_.Name = `Sunita "Suni" Williams`
	__Resource__00000002_.Description = ``
	__Resource__00000002_.IsExpanded = false
	__Resource__00000002_.ComputedPrefix = `1.1.1.2`
	__Resource__00000002_.IsInRenameMode = false

	__Resource__00000003_.Name = `NASA`
	__Resource__00000003_.Description = ``
	__Resource__00000003_.IsExpanded = false
	__Resource__00000003_.ComputedPrefix = `1`
	__Resource__00000003_.IsInRenameMode = false

	__Resource__00000004_.Name = `Crew Commercial Program (CPP)`
	__Resource__00000004_.Description = ``
	__Resource__00000004_.IsExpanded = false
	__Resource__00000004_.ComputedPrefix = `1.1`
	__Resource__00000004_.IsInRenameMode = false

	__Resource__00000005_.Name = `Crews`
	__Resource__00000005_.Description = ``
	__Resource__00000005_.IsExpanded = false
	__Resource__00000005_.ComputedPrefix = `1.1.1`
	__Resource__00000005_.IsInRenameMode = false

	__Resource__00000006_.Name = `Boeing`
	__Resource__00000006_.Description = ``
	__Resource__00000006_.IsExpanded = false
	__Resource__00000006_.ComputedPrefix = `2`
	__Resource__00000006_.IsInRenameMode = false

	__Resource__00000007_.Name = ``
	__Resource__00000007_.Description = ``
	__Resource__00000007_.IsExpanded = false
	__Resource__00000007_.ComputedPrefix = `2.1`
	__Resource__00000007_.IsInRenameMode = false

	__Resource__00000008_.Name = ``
	__Resource__00000008_.Description = ``
	__Resource__00000008_.IsExpanded = false
	__Resource__00000008_.ComputedPrefix = `1.3`
	__Resource__00000008_.IsInRenameMode = false

	__Resource__00000009_.Name = ` Starliner Tests and Anomalies Review (STAR) Investigation Team`
	__Resource__00000009_.Description = ``
	__Resource__00000009_.IsExpanded = false
	__Resource__00000009_.ComputedPrefix = `1.1.2`
	__Resource__00000009_.IsInRenameMode = false

	__ResourceCompositionShape__00000000_.Name = `NASA to `
	__ResourceCompositionShape__00000000_.StartRatio = 0.500000
	__ResourceCompositionShape__00000000_.EndRatio = 0.500000
	__ResourceCompositionShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000000_.CornerOffsetRatio = 1.680000

	__ResourceCompositionShape__00000001_.Name = `Crew Commercial Program (CPP) to `
	__ResourceCompositionShape__00000001_.StartRatio = 0.500000
	__ResourceCompositionShape__00000001_.EndRatio = 0.500000
	__ResourceCompositionShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000001_.CornerOffsetRatio = 1.680000

	__ResourceCompositionShape__00000002_.Name = `Crews to Barry "Butch" Wilmore`
	__ResourceCompositionShape__00000002_.StartRatio = 0.500000
	__ResourceCompositionShape__00000002_.EndRatio = 0.500000
	__ResourceCompositionShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000002_.CornerOffsetRatio = 1.680000

	__ResourceCompositionShape__00000003_.Name = `Crews to Sunita "Suni" Williams`
	__ResourceCompositionShape__00000003_.StartRatio = 0.500000
	__ResourceCompositionShape__00000003_.EndRatio = 0.500000
	__ResourceCompositionShape__00000003_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000003_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000003_.CornerOffsetRatio = 1.680000

	__ResourceCompositionShape__00000004_.Name = `NASA to PITProgram Investigation Team (PIT)`
	__ResourceCompositionShape__00000004_.StartRatio = 0.500000
	__ResourceCompositionShape__00000004_.EndRatio = 0.500000
	__ResourceCompositionShape__00000004_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000004_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000004_.CornerOffsetRatio = 1.680000

	__ResourceCompositionShape__00000005_.Name = `Boeing to `
	__ResourceCompositionShape__00000005_.StartRatio = 0.500000
	__ResourceCompositionShape__00000005_.EndRatio = 0.500000
	__ResourceCompositionShape__00000005_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000005_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000005_.CornerOffsetRatio = 1.680000

	__ResourceCompositionShape__00000006_.Name = `NASA to `
	__ResourceCompositionShape__00000006_.StartRatio = 0.500000
	__ResourceCompositionShape__00000006_.EndRatio = 0.500000
	__ResourceCompositionShape__00000006_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000006_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000006_.CornerOffsetRatio = 1.680000

	__ResourceCompositionShape__00000007_.Name = `Crew Commercial Program (CPP) to `
	__ResourceCompositionShape__00000007_.StartRatio = 0.500000
	__ResourceCompositionShape__00000007_.EndRatio = 0.500000
	__ResourceCompositionShape__00000007_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000007_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000007_.CornerOffsetRatio = 1.680000

	__ResourceShape__00000000_.Name = `-Default Diagram`
	__ResourceShape__00000000_.IsExpanded = false
	__ResourceShape__00000000_.X = 52.114853
	__ResourceShape__00000000_.Y = 34.127119
	__ResourceShape__00000000_.Width = 250.000000
	__ResourceShape__00000000_.Height = 70.000000

	__ResourceShape__00000001_.Name = `-Default Diagram`
	__ResourceShape__00000001_.IsExpanded = false
	__ResourceShape__00000001_.X = 34.180389
	__ResourceShape__00000001_.Y = 553.010316
	__ResourceShape__00000001_.Width = 250.000000
	__ResourceShape__00000001_.Height = 70.000000

	__ResourceShape__00000002_.Name = `-Default Diagram`
	__ResourceShape__00000002_.IsExpanded = false
	__ResourceShape__00000002_.X = 33.201744
	__ResourceShape__00000002_.Y = 650.719328
	__ResourceShape__00000002_.Width = 250.000000
	__ResourceShape__00000002_.Height = 70.000000

	__ResourceShape__00000003_.Name = `-RBS`
	__ResourceShape__00000003_.IsExpanded = false
	__ResourceShape__00000003_.X = 134.737820
	__ResourceShape__00000003_.Y = 123.124371
	__ResourceShape__00000003_.Width = 250.000000
	__ResourceShape__00000003_.Height = 70.000000

	__ResourceShape__00000004_.Name = `-RBS`
	__ResourceShape__00000004_.IsExpanded = false
	__ResourceShape__00000004_.X = 137.737820
	__ResourceShape__00000004_.Y = 300.124371
	__ResourceShape__00000004_.Width = 250.000000
	__ResourceShape__00000004_.Height = 70.000000

	__ResourceShape__00000005_.Name = `-RBS`
	__ResourceShape__00000005_.IsExpanded = false
	__ResourceShape__00000005_.X = 137.737820
	__ResourceShape__00000005_.Y = 440.124371
	__ResourceShape__00000005_.Width = 250.000000
	__ResourceShape__00000005_.Height = 70.000000

	__ResourceShape__00000006_.Name = `Barry "Butch" Wilmore-RBS`
	__ResourceShape__00000006_.IsExpanded = false
	__ResourceShape__00000006_.X = 132.423066
	__ResourceShape__00000006_.Y = 597.719206
	__ResourceShape__00000006_.Width = 250.000000
	__ResourceShape__00000006_.Height = 70.000000

	__ResourceShape__00000007_.Name = `Sunita "Suni" Williams-RBS`
	__ResourceShape__00000007_.IsExpanded = false
	__ResourceShape__00000007_.X = 454.596590
	__ResourceShape__00000007_.Y = 600.173621
	__ResourceShape__00000007_.Width = 250.000000
	__ResourceShape__00000007_.Height = 70.000000

	__ResourceShape__00000008_.Name = `PITProgram Investigation Team (PIT)-RBS`
	__ResourceShape__00000008_.IsExpanded = false
	__ResourceShape__00000008_.X = 835.651635
	__ResourceShape__00000008_.Y = 297.886860
	__ResourceShape__00000008_.Width = 250.000000
	__ResourceShape__00000008_.Height = 70.000000

	__ResourceShape__00000009_.Name = `-RBS`
	__ResourceShape__00000009_.IsExpanded = false
	__ResourceShape__00000009_.X = 126.434105
	__ResourceShape__00000009_.Y = 846.201664
	__ResourceShape__00000009_.Width = 250.000000
	__ResourceShape__00000009_.Height = 70.000000

	__ResourceShape__00000010_.Name = `-RBS`
	__ResourceShape__00000010_.IsExpanded = false
	__ResourceShape__00000010_.X = 110.434105
	__ResourceShape__00000010_.Y = 1029.201740
	__ResourceShape__00000010_.Width = 250.000000
	__ResourceShape__00000010_.Height = 70.000000

	__ResourceShape__00000011_.Name = `-RBS`
	__ResourceShape__00000011_.IsExpanded = false
	__ResourceShape__00000011_.X = 136.986696
	__ResourceShape__00000011_.Y = 123.307852
	__ResourceShape__00000011_.Width = 250.000000
	__ResourceShape__00000011_.Height = 70.000000

	__ResourceShape__00000012_.Name = `-RBS`
	__ResourceShape__00000012_.IsExpanded = false
	__ResourceShape__00000012_.X = 437.737820
	__ResourceShape__00000012_.Y = 440.124371
	__ResourceShape__00000012_.Width = 250.000000
	__ResourceShape__00000012_.Height = 70.000000

	__ResourceShape__00000013_.Name = `Program Investigation Team (PIT)-PBS`
	__ResourceShape__00000013_.IsExpanded = false
	__ResourceShape__00000013_.X = 1128.684052
	__ResourceShape__00000013_.Y = 492.417667
	__ResourceShape__00000013_.Width = 250.000000
	__ResourceShape__00000013_.Height = 70.000000

	__ResourceShape__00000014_.Name = `Program Investigation Team (PIT)-PIT focus`
	__ResourceShape__00000014_.IsExpanded = false
	__ResourceShape__00000014_.X = 104.267926
	__ResourceShape__00000014_.Y = 288.658205
	__ResourceShape__00000014_.Width = 250.000000
	__ResourceShape__00000014_.Height = 70.000000

	__ResourceTaskShape__00000000_.Name = `PITProgram Investigation Team (PIT) to Mishap investigation`
	__ResourceTaskShape__00000000_.StartRatio = 0.500000
	__ResourceTaskShape__00000000_.EndRatio = 0.500000
	__ResourceTaskShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceTaskShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceTaskShape__00000000_.CornerOffsetRatio = 1.680000

	__Root__00000000_.Name = `Root`
	__Root__00000000_.NbPixPerCharacter = 8.000000

	__Task__00000000_.Name = `Starliner Crewed Flight Test (CFT)`
	__Task__00000000_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2024-06-05 00:00:00 +0000 UTC")
	__Task__00000000_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2024-06-10 00:00:00 +0000 UTC")
	__Task__00000000_.Description = `The mission of interest to the report.`
	__Task__00000000_.IsExpanded = false
	__Task__00000000_.ComputedPrefix = `2.3`
	__Task__00000000_.IsInRenameMode = false
	__Task__00000000_.IsInputsNodeExpanded = false
	__Task__00000000_.IsOutputsNodeExpanded = false
	__Task__00000000_.IsWithCompletion = false
	__Task__00000000_.Completion = ""

	__Task__00000001_.Name = `Mishap investigations`
	__Task__00000001_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2025-02-01 00:00:00 +0000 UTC")
	__Task__00000001_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2025-11-01 00:00:00 +0000 UTC")
	__Task__00000001_.Description = ``
	__Task__00000001_.IsExpanded = false
	__Task__00000001_.ComputedPrefix = `1`
	__Task__00000001_.IsInRenameMode = false
	__Task__00000001_.IsInputsNodeExpanded = false
	__Task__00000001_.IsOutputsNodeExpanded = false
	__Task__00000001_.IsWithCompletion = false
	__Task__00000001_.Completion = ""

	__Task__00000002_.Name = `Commercial Crew Program (CCP),`
	__Task__00000002_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000002_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000002_.Description = ``
	__Task__00000002_.IsExpanded = false
	__Task__00000002_.ComputedPrefix = `2`
	__Task__00000002_.IsInRenameMode = false
	__Task__00000002_.IsInputsNodeExpanded = false
	__Task__00000002_.IsOutputsNodeExpanded = false
	__Task__00000002_.IsWithCompletion = false
	__Task__00000002_.Completion = ""

	__Task__00000003_.Name = ` Commercial ReSupply (CRS) `
	__Task__00000003_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000003_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000003_.Description = ``
	__Task__00000003_.IsExpanded = false
	__Task__00000003_.ComputedPrefix = `3`
	__Task__00000003_.IsInRenameMode = false
	__Task__00000003_.IsInputsNodeExpanded = false
	__Task__00000003_.IsOutputsNodeExpanded = false
	__Task__00000003_.IsWithCompletion = false
	__Task__00000003_.Completion = ""

	__Task__00000004_.Name = `Program Investigation Team (PIT) Report`
	__Task__00000004_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000004_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000004_.Description = ``
	__Task__00000004_.IsExpanded = false
	__Task__00000004_.ComputedPrefix = `1.1`
	__Task__00000004_.IsInRenameMode = false
	__Task__00000004_.IsInputsNodeExpanded = false
	__Task__00000004_.IsOutputsNodeExpanded = false
	__Task__00000004_.IsWithCompletion = false
	__Task__00000004_.Completion = ""

	__Task__00000005_.Name = `Orbital Flight Test-1 (OFT-1)`
	__Task__00000005_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2019-12-20 00:00:00 +0000 UTC")
	__Task__00000005_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2019-12-20 00:00:00 +0000 UTC")
	__Task__00000005_.Description = ``
	__Task__00000005_.IsExpanded = false
	__Task__00000005_.ComputedPrefix = `2.1`
	__Task__00000005_.IsInRenameMode = false
	__Task__00000005_.IsInputsNodeExpanded = false
	__Task__00000005_.IsOutputsNodeExpanded = false
	__Task__00000005_.IsWithCompletion = false
	__Task__00000005_.Completion = ""

	__Task__00000006_.Name = `Orbital Flight Test-2 (OFT-2)`
	__Task__00000006_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2022-05-19 00:00:00 +0000 UTC")
	__Task__00000006_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2022-05-25 00:00:00 +0000 UTC")
	__Task__00000006_.Description = ``
	__Task__00000006_.IsExpanded = false
	__Task__00000006_.ComputedPrefix = `2.2`
	__Task__00000006_.IsInRenameMode = false
	__Task__00000006_.IsInputsNodeExpanded = false
	__Task__00000006_.IsOutputsNodeExpanded = false
	__Task__00000006_.IsWithCompletion = false
	__Task__00000006_.Completion = ""

	__TaskCompositionShape__00000000_.Name = `Commercial Crew Program (CCP), to Starliner Crewed Flight Test (CFT)`
	__TaskCompositionShape__00000000_.StartRatio = 0.500000
	__TaskCompositionShape__00000000_.EndRatio = 0.500000
	__TaskCompositionShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000000_.CornerOffsetRatio = 1.680000

	__TaskCompositionShape__00000001_.Name = `Commercial Crew Program (CCP), to Starliner Crewed Flight Test (CFT)`
	__TaskCompositionShape__00000001_.StartRatio = 0.500000
	__TaskCompositionShape__00000001_.EndRatio = 0.500000
	__TaskCompositionShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000001_.CornerOffsetRatio = 1.680000

	__TaskCompositionShape__00000002_.Name = `Mishap investigations to `
	__TaskCompositionShape__00000002_.StartRatio = 0.500000
	__TaskCompositionShape__00000002_.EndRatio = 0.500000
	__TaskCompositionShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000002_.CornerOffsetRatio = 1.680000

	__TaskCompositionShape__00000003_.Name = `Commercial Crew Program (CCP), to `
	__TaskCompositionShape__00000003_.StartRatio = 0.500000
	__TaskCompositionShape__00000003_.EndRatio = 0.500000
	__TaskCompositionShape__00000003_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000003_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000003_.CornerOffsetRatio = 1.680000

	__TaskCompositionShape__00000004_.Name = `Commercial Crew Program (CCP), to `
	__TaskCompositionShape__00000004_.StartRatio = 0.500000
	__TaskCompositionShape__00000004_.EndRatio = 0.500000
	__TaskCompositionShape__00000004_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000004_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000004_.CornerOffsetRatio = 1.680000

	__TaskInputShape__00000000_.Name = `Program Investigation Team (PIT) Report to  Starliner Tests and Anomalies Review (STAR) Investigation Report`
	__TaskInputShape__00000000_.StartRatio = 0.278248
	__TaskInputShape__00000000_.EndRatio = 0.277627
	__TaskInputShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskInputShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskInputShape__00000000_.CornerOffsetRatio = 1.883372

	__TaskOutputShape__00000000_.Name = `Program Investigation Team (PIT) Report to Program Investigation Team (PIT) Report`
	__TaskOutputShape__00000000_.StartRatio = 0.439622
	__TaskOutputShape__00000000_.EndRatio = 0.501995
	__TaskOutputShape__00000000_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000000_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000000_.CornerOffsetRatio = 1.313627

	__TaskShape__00000000_.Name = `NewTask-Default Diagram`
	__TaskShape__00000000_.IsExpanded = false
	__TaskShape__00000000_.X = 61.409529
	__TaskShape__00000000_.Y = 952.589880
	__TaskShape__00000000_.Width = 250.000000
	__TaskShape__00000000_.Height = 70.000000

	__TaskShape__00000001_.Name = `NewTask-Default Diagram`
	__TaskShape__00000001_.IsExpanded = false
	__TaskShape__00000001_.X = 105.187127
	__TaskShape__00000001_.Y = 246.958864
	__TaskShape__00000001_.Width = 250.000000
	__TaskShape__00000001_.Height = 70.000000

	__TaskShape__00000002_.Name = `-Default Diagram`
	__TaskShape__00000002_.IsExpanded = false
	__TaskShape__00000002_.X = 62.073018
	__TaskShape__00000002_.Y = 795.651222
	__TaskShape__00000002_.Width = 250.000000
	__TaskShape__00000002_.Height = 70.000000

	__TaskShape__00000003_.Name = `Mishap investigation-WBS`
	__TaskShape__00000003_.IsExpanded = false
	__TaskShape__00000003_.X = 28.416829
	__TaskShape__00000003_.Y = 546.674065
	__TaskShape__00000003_.Width = 250.000000
	__TaskShape__00000003_.Height = 70.000000

	__TaskShape__00000004_.Name = `Commercial Crew Program (CCP),-WBS`
	__TaskShape__00000004_.IsExpanded = false
	__TaskShape__00000004_.X = 16.071983
	__TaskShape__00000004_.Y = 166.517160
	__TaskShape__00000004_.Width = 250.000000
	__TaskShape__00000004_.Height = 70.000000

	__TaskShape__00000005_.Name = `Starliner Crewed Flight Test (CFT)-WBS`
	__TaskShape__00000005_.IsExpanded = false
	__TaskShape__00000005_.X = 645.617766
	__TaskShape__00000005_.Y = 360.295606
	__TaskShape__00000005_.Width = 250.000000
	__TaskShape__00000005_.Height = 70.000000

	__TaskShape__00000006_.Name = `-WBS`
	__TaskShape__00000006_.IsExpanded = false
	__TaskShape__00000006_.X = 14.409715
	__TaskShape__00000006_.Y = 32.470060
	__TaskShape__00000006_.Width = 250.000000
	__TaskShape__00000006_.Height = 70.000000

	__TaskShape__00000007_.Name = `Mishap investigation-PBS`
	__TaskShape__00000007_.IsExpanded = false
	__TaskShape__00000007_.X = 1128.178322
	__TaskShape__00000007_.Y = 381.328484
	__TaskShape__00000007_.Width = 250.000000
	__TaskShape__00000007_.Height = 70.000000

	__TaskShape__00000008_.Name = `-WBS`
	__TaskShape__00000008_.IsExpanded = false
	__TaskShape__00000008_.X = 28.416829
	__TaskShape__00000008_.Y = 686.674065
	__TaskShape__00000008_.Width = 250.000000
	__TaskShape__00000008_.Height = 70.000000

	__TaskShape__00000009_.Name = `Program Investigation Team (PIT) Report-PIT focus`
	__TaskShape__00000009_.IsExpanded = false
	__TaskShape__00000009_.X = 499.174738
	__TaskShape__00000009_.Y = 283.233412
	__TaskShape__00000009_.Width = 250.000000
	__TaskShape__00000009_.Height = 70.000000

	__TaskShape__00000010_.Name = `-WBS`
	__TaskShape__00000010_.IsExpanded = false
	__TaskShape__00000010_.X = 19.071983
	__TaskShape__00000010_.Y = 346.517160
	__TaskShape__00000010_.Width = 250.000000
	__TaskShape__00000010_.Height = 70.000000

	__TaskShape__00000011_.Name = `-WBS`
	__TaskShape__00000011_.IsExpanded = false
	__TaskShape__00000011_.X = 336.071983
	__TaskShape__00000011_.Y = 355.517160
	__TaskShape__00000011_.Width = 250.000000
	__TaskShape__00000011_.Height = 70.000000

	// insertion point for setup of pointers
	__Diagram__00000000_.Task_Shapes = append(__Diagram__00000000_.Task_Shapes, __TaskShape__00000000_)
	__Diagram__00000000_.Task_Shapes = append(__Diagram__00000000_.Task_Shapes, __TaskShape__00000001_)
	__Diagram__00000000_.Task_Shapes = append(__Diagram__00000000_.Task_Shapes, __TaskShape__00000002_)
	__Diagram__00000000_.TasksWhoseNodeIsExpanded = append(__Diagram__00000000_.TasksWhoseNodeIsExpanded, __Task__00000002_)
	__Diagram__00000000_.TaskComposition_Shapes = append(__Diagram__00000000_.TaskComposition_Shapes, __TaskCompositionShape__00000000_)
	__Diagram__00000000_.Note_Shapes = append(__Diagram__00000000_.Note_Shapes, __NoteShape__00000000_)
	__Diagram__00000000_.NoteTaskShapes = append(__Diagram__00000000_.NoteTaskShapes, __NoteTaskShape__00000000_)
	__Diagram__00000000_.Resource_Shapes = append(__Diagram__00000000_.Resource_Shapes, __ResourceShape__00000000_)
	__Diagram__00000000_.Resource_Shapes = append(__Diagram__00000000_.Resource_Shapes, __ResourceShape__00000001_)
	__Diagram__00000000_.Resource_Shapes = append(__Diagram__00000000_.Resource_Shapes, __ResourceShape__00000002_)
	__Diagram__00000000_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000000_.ResourcesWhoseNodeIsExpanded, __Resource__00000000_)
	__Diagram__00000000_.ResourceTaskShapes = append(__Diagram__00000000_.ResourceTaskShapes, __ResourceTaskShape__00000000_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000003_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000004_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000005_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000006_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000008_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000010_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000011_)
	__Diagram__00000001_.TasksWhoseNodeIsExpanded = append(__Diagram__00000001_.TasksWhoseNodeIsExpanded, __Task__00000002_)
	__Diagram__00000001_.TasksWhoseNodeIsExpanded = append(__Diagram__00000001_.TasksWhoseNodeIsExpanded, __Task__00000001_)
	__Diagram__00000001_.TaskComposition_Shapes = append(__Diagram__00000001_.TaskComposition_Shapes, __TaskCompositionShape__00000001_)
	__Diagram__00000001_.TaskComposition_Shapes = append(__Diagram__00000001_.TaskComposition_Shapes, __TaskCompositionShape__00000002_)
	__Diagram__00000001_.TaskComposition_Shapes = append(__Diagram__00000001_.TaskComposition_Shapes, __TaskCompositionShape__00000003_)
	__Diagram__00000001_.TaskComposition_Shapes = append(__Diagram__00000001_.TaskComposition_Shapes, __TaskCompositionShape__00000004_)
	__Diagram__00000001_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000001_.ResourcesWhoseNodeIsExpanded, __Resource__00000003_)
	__Diagram__00000002_.Product_Shapes = append(__Diagram__00000002_.Product_Shapes, __ProductShape__00000002_)
	__Diagram__00000002_.Product_Shapes = append(__Diagram__00000002_.Product_Shapes, __ProductShape__00000003_)
	__Diagram__00000002_.Product_Shapes = append(__Diagram__00000002_.Product_Shapes, __ProductShape__00000004_)
	__Diagram__00000002_.Product_Shapes = append(__Diagram__00000002_.Product_Shapes, __ProductShape__00000005_)
	__Diagram__00000002_.Product_Shapes = append(__Diagram__00000002_.Product_Shapes, __ProductShape__00000006_)
	__Diagram__00000002_.Product_Shapes = append(__Diagram__00000002_.Product_Shapes, __ProductShape__00000009_)
	__Diagram__00000002_.Product_Shapes = append(__Diagram__00000002_.Product_Shapes, __ProductShape__00000010_)
	__Diagram__00000002_.Product_Shapes = append(__Diagram__00000002_.Product_Shapes, __ProductShape__00000011_)
	__Diagram__00000002_.Product_Shapes = append(__Diagram__00000002_.Product_Shapes, __ProductShape__00000027_)
	__Diagram__00000002_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000002_.ProductsWhoseNodeIsExpanded, __Product__00000003_)
	__Diagram__00000002_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000002_.ProductsWhoseNodeIsExpanded, __Product__00000006_)
	__Diagram__00000002_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000002_.ProductsWhoseNodeIsExpanded, __Product__00000007_)
	__Diagram__00000002_.ProductComposition_Shapes = append(__Diagram__00000002_.ProductComposition_Shapes, __ProductCompositionShape__00000000_)
	__Diagram__00000002_.ProductComposition_Shapes = append(__Diagram__00000002_.ProductComposition_Shapes, __ProductCompositionShape__00000001_)
	__Diagram__00000002_.ProductComposition_Shapes = append(__Diagram__00000002_.ProductComposition_Shapes, __ProductCompositionShape__00000002_)
	__Diagram__00000002_.ProductComposition_Shapes = append(__Diagram__00000002_.ProductComposition_Shapes, __ProductCompositionShape__00000003_)
	__Diagram__00000002_.ProductComposition_Shapes = append(__Diagram__00000002_.ProductComposition_Shapes, __ProductCompositionShape__00000004_)
	__Diagram__00000002_.ProductComposition_Shapes = append(__Diagram__00000002_.ProductComposition_Shapes, __ProductCompositionShape__00000005_)
	__Diagram__00000002_.ProductComposition_Shapes = append(__Diagram__00000002_.ProductComposition_Shapes, __ProductCompositionShape__00000018_)
	__Diagram__00000002_.Task_Shapes = append(__Diagram__00000002_.Task_Shapes, __TaskShape__00000007_)
	__Diagram__00000002_.TasksWhoseNodeIsExpanded = append(__Diagram__00000002_.TasksWhoseNodeIsExpanded, __Task__00000002_)
	__Diagram__00000002_.Note_Shapes = append(__Diagram__00000002_.Note_Shapes, __NoteShape__00000002_)
	__Diagram__00000002_.NoteProductShapes = append(__Diagram__00000002_.NoteProductShapes, __NoteProductShape__00000001_)
	__Diagram__00000002_.Resource_Shapes = append(__Diagram__00000002_.Resource_Shapes, __ResourceShape__00000013_)
	__Diagram__00000002_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000002_.ResourcesWhoseNodeIsExpanded, __Resource__00000003_)
	__Diagram__00000002_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000002_.ResourcesWhoseNodeIsExpanded, __Resource__00000000_)
	__Diagram__00000002_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000002_.ResourcesWhoseNodeIsExpanded, __Resource__00000006_)
	__Diagram__00000003_.Resource_Shapes = append(__Diagram__00000003_.Resource_Shapes, __ResourceShape__00000003_)
	__Diagram__00000003_.Resource_Shapes = append(__Diagram__00000003_.Resource_Shapes, __ResourceShape__00000004_)
	__Diagram__00000003_.Resource_Shapes = append(__Diagram__00000003_.Resource_Shapes, __ResourceShape__00000005_)
	__Diagram__00000003_.Resource_Shapes = append(__Diagram__00000003_.Resource_Shapes, __ResourceShape__00000006_)
	__Diagram__00000003_.Resource_Shapes = append(__Diagram__00000003_.Resource_Shapes, __ResourceShape__00000007_)
	__Diagram__00000003_.Resource_Shapes = append(__Diagram__00000003_.Resource_Shapes, __ResourceShape__00000008_)
	__Diagram__00000003_.Resource_Shapes = append(__Diagram__00000003_.Resource_Shapes, __ResourceShape__00000009_)
	__Diagram__00000003_.Resource_Shapes = append(__Diagram__00000003_.Resource_Shapes, __ResourceShape__00000010_)
	__Diagram__00000003_.Resource_Shapes = append(__Diagram__00000003_.Resource_Shapes, __ResourceShape__00000011_)
	__Diagram__00000003_.Resource_Shapes = append(__Diagram__00000003_.Resource_Shapes, __ResourceShape__00000012_)
	__Diagram__00000003_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000003_.ResourcesWhoseNodeIsExpanded, __Resource__00000000_)
	__Diagram__00000003_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000003_.ResourcesWhoseNodeIsExpanded, __Resource__00000003_)
	__Diagram__00000003_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000003_.ResourcesWhoseNodeIsExpanded, __Resource__00000004_)
	__Diagram__00000003_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000003_.ResourcesWhoseNodeIsExpanded, __Resource__00000006_)
	__Diagram__00000003_.ResourceComposition_Shapes = append(__Diagram__00000003_.ResourceComposition_Shapes, __ResourceCompositionShape__00000000_)
	__Diagram__00000003_.ResourceComposition_Shapes = append(__Diagram__00000003_.ResourceComposition_Shapes, __ResourceCompositionShape__00000001_)
	__Diagram__00000003_.ResourceComposition_Shapes = append(__Diagram__00000003_.ResourceComposition_Shapes, __ResourceCompositionShape__00000002_)
	__Diagram__00000003_.ResourceComposition_Shapes = append(__Diagram__00000003_.ResourceComposition_Shapes, __ResourceCompositionShape__00000003_)
	__Diagram__00000003_.ResourceComposition_Shapes = append(__Diagram__00000003_.ResourceComposition_Shapes, __ResourceCompositionShape__00000004_)
	__Diagram__00000003_.ResourceComposition_Shapes = append(__Diagram__00000003_.ResourceComposition_Shapes, __ResourceCompositionShape__00000005_)
	__Diagram__00000003_.ResourceComposition_Shapes = append(__Diagram__00000003_.ResourceComposition_Shapes, __ResourceCompositionShape__00000006_)
	__Diagram__00000003_.ResourceComposition_Shapes = append(__Diagram__00000003_.ResourceComposition_Shapes, __ResourceCompositionShape__00000007_)
	__Diagram__00000004_.Product_Shapes = append(__Diagram__00000004_.Product_Shapes, __ProductShape__00000007_)
	__Diagram__00000004_.Product_Shapes = append(__Diagram__00000004_.Product_Shapes, __ProductShape__00000008_)
	__Diagram__00000004_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000004_.ProductsWhoseNodeIsExpanded, __Product__00000003_)
	__Diagram__00000004_.Task_Shapes = append(__Diagram__00000004_.Task_Shapes, __TaskShape__00000009_)
	__Diagram__00000004_.TasksWhoseNodeIsExpanded = append(__Diagram__00000004_.TasksWhoseNodeIsExpanded, __Task__00000001_)
	__Diagram__00000004_.TasksWhoseNodeIsExpanded = append(__Diagram__00000004_.TasksWhoseNodeIsExpanded, __Task__00000004_)
	__Diagram__00000004_.TasksWhoseInputNodeIsExpanded = append(__Diagram__00000004_.TasksWhoseInputNodeIsExpanded, __Task__00000004_)
	__Diagram__00000004_.TasksWhoseOutputNodeIsExpanded = append(__Diagram__00000004_.TasksWhoseOutputNodeIsExpanded, __Task__00000004_)
	__Diagram__00000004_.TaskInputShapes = append(__Diagram__00000004_.TaskInputShapes, __TaskInputShape__00000000_)
	__Diagram__00000004_.TaskOutputShapes = append(__Diagram__00000004_.TaskOutputShapes, __TaskOutputShape__00000000_)
	__Diagram__00000004_.Note_Shapes = append(__Diagram__00000004_.Note_Shapes, __NoteShape__00000001_)
	__Diagram__00000004_.NoteProductShapes = append(__Diagram__00000004_.NoteProductShapes, __NoteProductShape__00000000_)
	__Diagram__00000004_.Resource_Shapes = append(__Diagram__00000004_.Resource_Shapes, __ResourceShape__00000014_)
	__Diagram__00000004_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000004_.ResourcesWhoseNodeIsExpanded, __Resource__00000003_)
	__Diagram__00000004_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000004_.ResourcesWhoseNodeIsExpanded, __Resource__00000000_)
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000012_)
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000013_)
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000014_)
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000022_)
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000023_)
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000028_)
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000029_)
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000032_)
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000034_)
	__Diagram__00000005_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000005_.ProductsWhoseNodeIsExpanded, __Product__00000003_)
	__Diagram__00000005_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000005_.ProductsWhoseNodeIsExpanded, __Product__00000005_)
	__Diagram__00000005_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000005_.ProductsWhoseNodeIsExpanded, __Product__00000007_)
	__Diagram__00000005_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000005_.ProductsWhoseNodeIsExpanded, __Product__00000006_)
	__Diagram__00000005_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000005_.ProductsWhoseNodeIsExpanded, __Product__00000017_)
	__Diagram__00000005_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000005_.ProductsWhoseNodeIsExpanded, __Product__00000009_)
	__Diagram__00000005_.ProductComposition_Shapes = append(__Diagram__00000005_.ProductComposition_Shapes, __ProductCompositionShape__00000006_)
	__Diagram__00000005_.ProductComposition_Shapes = append(__Diagram__00000005_.ProductComposition_Shapes, __ProductCompositionShape__00000013_)
	__Diagram__00000005_.ProductComposition_Shapes = append(__Diagram__00000005_.ProductComposition_Shapes, __ProductCompositionShape__00000014_)
	__Diagram__00000005_.ProductComposition_Shapes = append(__Diagram__00000005_.ProductComposition_Shapes, __ProductCompositionShape__00000019_)
	__Diagram__00000005_.ProductComposition_Shapes = append(__Diagram__00000005_.ProductComposition_Shapes, __ProductCompositionShape__00000020_)
	__Diagram__00000005_.ProductComposition_Shapes = append(__Diagram__00000005_.ProductComposition_Shapes, __ProductCompositionShape__00000023_)
	__Diagram__00000005_.ProductComposition_Shapes = append(__Diagram__00000005_.ProductComposition_Shapes, __ProductCompositionShape__00000025_)
	__Diagram__00000005_.Note_Shapes = append(__Diagram__00000005_.Note_Shapes, __NoteShape__00000003_)
	__Diagram__00000005_.Note_Shapes = append(__Diagram__00000005_.Note_Shapes, __NoteShape__00000004_)
	__Diagram__00000005_.NoteProductShapes = append(__Diagram__00000005_.NoteProductShapes, __NoteProductShape__00000002_)
	__Diagram__00000005_.NoteProductShapes = append(__Diagram__00000005_.NoteProductShapes, __NoteProductShape__00000003_)
	__Diagram__00000006_.Product_Shapes = append(__Diagram__00000006_.Product_Shapes, __ProductShape__00000015_)
	__Diagram__00000006_.Product_Shapes = append(__Diagram__00000006_.Product_Shapes, __ProductShape__00000016_)
	__Diagram__00000006_.Product_Shapes = append(__Diagram__00000006_.Product_Shapes, __ProductShape__00000017_)
	__Diagram__00000006_.Product_Shapes = append(__Diagram__00000006_.Product_Shapes, __ProductShape__00000018_)
	__Diagram__00000006_.Product_Shapes = append(__Diagram__00000006_.Product_Shapes, __ProductShape__00000019_)
	__Diagram__00000006_.Product_Shapes = append(__Diagram__00000006_.Product_Shapes, __ProductShape__00000020_)
	__Diagram__00000006_.Product_Shapes = append(__Diagram__00000006_.Product_Shapes, __ProductShape__00000021_)
	__Diagram__00000006_.Product_Shapes = append(__Diagram__00000006_.Product_Shapes, __ProductShape__00000024_)
	__Diagram__00000006_.Product_Shapes = append(__Diagram__00000006_.Product_Shapes, __ProductShape__00000025_)
	__Diagram__00000006_.Product_Shapes = append(__Diagram__00000006_.Product_Shapes, __ProductShape__00000026_)
	__Diagram__00000006_.Product_Shapes = append(__Diagram__00000006_.Product_Shapes, __ProductShape__00000030_)
	__Diagram__00000006_.Product_Shapes = append(__Diagram__00000006_.Product_Shapes, __ProductShape__00000031_)
	__Diagram__00000006_.Product_Shapes = append(__Diagram__00000006_.Product_Shapes, __ProductShape__00000033_)
	__Diagram__00000006_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000006_.ProductsWhoseNodeIsExpanded, __Product__00000006_)
	__Diagram__00000006_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000006_.ProductsWhoseNodeIsExpanded, __Product__00000012_)
	__Diagram__00000006_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000006_.ProductsWhoseNodeIsExpanded, __Product__00000002_)
	__Diagram__00000006_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000006_.ProductsWhoseNodeIsExpanded, __Product__00000014_)
	__Diagram__00000006_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000006_.ProductsWhoseNodeIsExpanded, __Product__00000015_)
	__Diagram__00000006_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000006_.ProductsWhoseNodeIsExpanded, __Product__00000011_)
	__Diagram__00000006_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000006_.ProductsWhoseNodeIsExpanded, __Product__00000007_)
	__Diagram__00000006_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000006_.ProductsWhoseNodeIsExpanded, __Product__00000013_)
	__Diagram__00000006_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000006_.ProductsWhoseNodeIsExpanded, __Product__00000003_)
	__Diagram__00000006_.ProductComposition_Shapes = append(__Diagram__00000006_.ProductComposition_Shapes, __ProductCompositionShape__00000007_)
	__Diagram__00000006_.ProductComposition_Shapes = append(__Diagram__00000006_.ProductComposition_Shapes, __ProductCompositionShape__00000008_)
	__Diagram__00000006_.ProductComposition_Shapes = append(__Diagram__00000006_.ProductComposition_Shapes, __ProductCompositionShape__00000009_)
	__Diagram__00000006_.ProductComposition_Shapes = append(__Diagram__00000006_.ProductComposition_Shapes, __ProductCompositionShape__00000010_)
	__Diagram__00000006_.ProductComposition_Shapes = append(__Diagram__00000006_.ProductComposition_Shapes, __ProductCompositionShape__00000011_)
	__Diagram__00000006_.ProductComposition_Shapes = append(__Diagram__00000006_.ProductComposition_Shapes, __ProductCompositionShape__00000012_)
	__Diagram__00000006_.ProductComposition_Shapes = append(__Diagram__00000006_.ProductComposition_Shapes, __ProductCompositionShape__00000015_)
	__Diagram__00000006_.ProductComposition_Shapes = append(__Diagram__00000006_.ProductComposition_Shapes, __ProductCompositionShape__00000016_)
	__Diagram__00000006_.ProductComposition_Shapes = append(__Diagram__00000006_.ProductComposition_Shapes, __ProductCompositionShape__00000017_)
	__Diagram__00000006_.ProductComposition_Shapes = append(__Diagram__00000006_.ProductComposition_Shapes, __ProductCompositionShape__00000021_)
	__Diagram__00000006_.ProductComposition_Shapes = append(__Diagram__00000006_.ProductComposition_Shapes, __ProductCompositionShape__00000022_)
	__Diagram__00000006_.ProductComposition_Shapes = append(__Diagram__00000006_.ProductComposition_Shapes, __ProductCompositionShape__00000024_)
	__Diagram__00000006_.Note_Shapes = append(__Diagram__00000006_.Note_Shapes, __NoteShape__00000005_)
	__Diagram__00000006_.Note_Shapes = append(__Diagram__00000006_.Note_Shapes, __NoteShape__00000006_)
	__Diagram__00000006_.Note_Shapes = append(__Diagram__00000006_.Note_Shapes, __NoteShape__00000007_)
	__Diagram__00000006_.Note_Shapes = append(__Diagram__00000006_.Note_Shapes, __NoteShape__00000008_)
	__Diagram__00000006_.Note_Shapes = append(__Diagram__00000006_.Note_Shapes, __NoteShape__00000009_)
	__Diagram__00000006_.Note_Shapes = append(__Diagram__00000006_.Note_Shapes, __NoteShape__00000010_)
	__Diagram__00000006_.Note_Shapes = append(__Diagram__00000006_.Note_Shapes, __NoteShape__00000011_)
	__Diagram__00000006_.Note_Shapes = append(__Diagram__00000006_.Note_Shapes, __NoteShape__00000012_)
	__Diagram__00000006_.NoteProductShapes = append(__Diagram__00000006_.NoteProductShapes, __NoteProductShape__00000004_)
	__Diagram__00000006_.NoteProductShapes = append(__Diagram__00000006_.NoteProductShapes, __NoteProductShape__00000005_)
	__Diagram__00000006_.NoteProductShapes = append(__Diagram__00000006_.NoteProductShapes, __NoteProductShape__00000006_)
	__Diagram__00000006_.NoteProductShapes = append(__Diagram__00000006_.NoteProductShapes, __NoteProductShape__00000007_)
	__Diagram__00000006_.NoteProductShapes = append(__Diagram__00000006_.NoteProductShapes, __NoteProductShape__00000008_)
	__Diagram__00000006_.NoteProductShapes = append(__Diagram__00000006_.NoteProductShapes, __NoteProductShape__00000009_)
	__Diagram__00000006_.NoteProductShapes = append(__Diagram__00000006_.NoteProductShapes, __NoteProductShape__00000010_)
	__Diagram__00000006_.NoteProductShapes = append(__Diagram__00000006_.NoteProductShapes, __NoteProductShape__00000011_)
	__Diagram__00000006_.NoteProductShapes = append(__Diagram__00000006_.NoteProductShapes, __NoteProductShape__00000012_)
	__Diagram__00000006_.NoteProductShapes = append(__Diagram__00000006_.NoteProductShapes, __NoteProductShape__00000013_)
	__Diagram__00000006_.NoteProductShapes = append(__Diagram__00000006_.NoteProductShapes, __NoteProductShape__00000014_)
	__Diagram__00000006_.NoteProductShapes = append(__Diagram__00000006_.NoteProductShapes, __NoteProductShape__00000015_)
	__Note__00000000_.Tasks = append(__Note__00000000_.Tasks, __Task__00000000_)
	__Note__00000001_.Products = append(__Note__00000001_.Products, __Product__00000004_)
	__Note__00000002_.Products = append(__Note__00000002_.Products, __Product__00000006_)
	__Note__00000003_.Products = append(__Note__00000003_.Products, __Product__00000010_)
	__Note__00000004_.Products = append(__Note__00000004_.Products, __Product__00000010_)
	__Note__00000005_.Products = append(__Note__00000005_.Products, __Product__00000013_)
	__Note__00000005_.Products = append(__Note__00000005_.Products, __Product__00000016_)
	__Note__00000006_.Products = append(__Note__00000006_.Products, __Product__00000016_)
	__Note__00000007_.Products = append(__Note__00000007_.Products, __Product__00000016_)
	__Note__00000008_.Products = append(__Note__00000008_.Products, __Product__00000019_)
	__Note__00000008_.Products = append(__Note__00000008_.Products, __Product__00000016_)
	__Note__00000008_.Products = append(__Note__00000008_.Products, __Product__00000013_)
	__Note__00000009_.Products = append(__Note__00000009_.Products, __Product__00000016_)
	__Note__00000010_.Products = append(__Note__00000010_.Products, __Product__00000016_)
	__Note__00000011_.Products = append(__Note__00000011_.Products, __Product__00000013_)
	__Note__00000011_.Products = append(__Note__00000011_.Products, __Product__00000028_)
	__Note__00000012_.Products = append(__Note__00000012_.Products, __Product__00000028_)
	__NoteProductShape__00000000_.Note = __Note__00000001_
	__NoteProductShape__00000000_.Product = __Product__00000004_
	__NoteProductShape__00000001_.Note = __Note__00000002_
	__NoteProductShape__00000001_.Product = __Product__00000006_
	__NoteProductShape__00000002_.Note = __Note__00000003_
	__NoteProductShape__00000002_.Product = __Product__00000010_
	__NoteProductShape__00000003_.Note = __Note__00000004_
	__NoteProductShape__00000003_.Product = __Product__00000010_
	__NoteProductShape__00000004_.Note = __Note__00000005_
	__NoteProductShape__00000004_.Product = __Product__00000013_
	__NoteProductShape__00000005_.Note = __Note__00000005_
	__NoteProductShape__00000005_.Product = __Product__00000016_
	__NoteProductShape__00000006_.Note = __Note__00000006_
	__NoteProductShape__00000006_.Product = __Product__00000016_
	__NoteProductShape__00000007_.Note = __Note__00000007_
	__NoteProductShape__00000007_.Product = __Product__00000016_
	__NoteProductShape__00000008_.Note = __Note__00000008_
	__NoteProductShape__00000008_.Product = __Product__00000019_
	__NoteProductShape__00000009_.Note = __Note__00000008_
	__NoteProductShape__00000009_.Product = __Product__00000016_
	__NoteProductShape__00000010_.Note = __Note__00000008_
	__NoteProductShape__00000010_.Product = __Product__00000013_
	__NoteProductShape__00000011_.Note = __Note__00000009_
	__NoteProductShape__00000011_.Product = __Product__00000016_
	__NoteProductShape__00000012_.Note = __Note__00000010_
	__NoteProductShape__00000012_.Product = __Product__00000016_
	__NoteProductShape__00000013_.Note = __Note__00000011_
	__NoteProductShape__00000013_.Product = __Product__00000013_
	__NoteProductShape__00000014_.Note = __Note__00000011_
	__NoteProductShape__00000014_.Product = __Product__00000028_
	__NoteProductShape__00000015_.Note = __Note__00000012_
	__NoteProductShape__00000015_.Product = __Product__00000028_
	__NoteShape__00000000_.Note = __Note__00000000_
	__NoteShape__00000001_.Note = __Note__00000001_
	__NoteShape__00000002_.Note = __Note__00000002_
	__NoteShape__00000003_.Note = __Note__00000003_
	__NoteShape__00000004_.Note = __Note__00000004_
	__NoteShape__00000005_.Note = __Note__00000005_
	__NoteShape__00000006_.Note = __Note__00000006_
	__NoteShape__00000007_.Note = __Note__00000007_
	__NoteShape__00000008_.Note = __Note__00000008_
	__NoteShape__00000009_.Note = __Note__00000009_
	__NoteShape__00000010_.Note = __Note__00000010_
	__NoteShape__00000011_.Note = __Note__00000011_
	__NoteShape__00000012_.Note = __Note__00000012_
	__NoteTaskShape__00000000_.Note = __Note__00000000_
	__NoteTaskShape__00000000_.Task = __Task__00000000_
	__Product__00000002_.SubProducts = append(__Product__00000002_.SubProducts, __Product__00000011_)
	__Product__00000002_.SubProducts = append(__Product__00000002_.SubProducts, __Product__00000014_)
	__Product__00000002_.SubProducts = append(__Product__00000002_.SubProducts, __Product__00000019_)
	__Product__00000002_.SubProducts = append(__Product__00000002_.SubProducts, __Product__00000021_)
	__Product__00000002_.SubProducts = append(__Product__00000002_.SubProducts, __Product__00000025_)
	__Product__00000002_.SubProducts = append(__Product__00000002_.SubProducts, __Product__00000026_)
	__Product__00000003_.SubProducts = append(__Product__00000003_.SubProducts, __Product__00000004_)
	__Product__00000003_.SubProducts = append(__Product__00000003_.SubProducts, __Product__00000005_)
	__Product__00000003_.SubProducts = append(__Product__00000003_.SubProducts, __Product__00000022_)
	__Product__00000005_.SubProducts = append(__Product__00000005_.SubProducts, __Product__00000017_)
	__Product__00000005_.SubProducts = append(__Product__00000005_.SubProducts, __Product__00000009_)
	__Product__00000005_.SubProducts = append(__Product__00000005_.SubProducts, __Product__00000029_)
	__Product__00000006_.SubProducts = append(__Product__00000006_.SubProducts, __Product__00000002_)
	__Product__00000006_.SubProducts = append(__Product__00000006_.SubProducts, __Product__00000001_)
	__Product__00000006_.SubProducts = append(__Product__00000006_.SubProducts, __Product__00000010_)
	__Product__00000007_.SubProducts = append(__Product__00000007_.SubProducts, __Product__00000006_)
	__Product__00000007_.SubProducts = append(__Product__00000007_.SubProducts, __Product__00000008_)
	__Product__00000009_.SubProducts = append(__Product__00000009_.SubProducts, __Product__00000023_)
	__Product__00000009_.SubProducts = append(__Product__00000009_.SubProducts, __Product__00000024_)
	__Product__00000009_.SubProducts = append(__Product__00000009_.SubProducts, __Product__00000027_)
	__Product__00000011_.SubProducts = append(__Product__00000011_.SubProducts, __Product__00000012_)
	__Product__00000012_.SubProducts = append(__Product__00000012_.SubProducts, __Product__00000013_)
	__Product__00000013_.SubProducts = append(__Product__00000013_.SubProducts, __Product__00000020_)
	__Product__00000013_.SubProducts = append(__Product__00000013_.SubProducts, __Product__00000028_)
	__Product__00000014_.SubProducts = append(__Product__00000014_.SubProducts, __Product__00000015_)
	__Product__00000015_.SubProducts = append(__Product__00000015_.SubProducts, __Product__00000016_)
	__Product__00000017_.SubProducts = append(__Product__00000017_.SubProducts, __Product__00000018_)
	__ProductCompositionShape__00000000_.Product = __Product__00000004_
	__ProductCompositionShape__00000001_.Product = __Product__00000005_
	__ProductCompositionShape__00000002_.Product = __Product__00000002_
	__ProductCompositionShape__00000003_.Product = __Product__00000001_
	__ProductCompositionShape__00000004_.Product = __Product__00000006_
	__ProductCompositionShape__00000005_.Product = __Product__00000008_
	__ProductCompositionShape__00000006_.Product = __Product__00000009_
	__ProductCompositionShape__00000007_.Product = __Product__00000012_
	__ProductCompositionShape__00000008_.Product = __Product__00000013_
	__ProductCompositionShape__00000009_.Product = __Product__00000011_
	__ProductCompositionShape__00000010_.Product = __Product__00000014_
	__ProductCompositionShape__00000011_.Product = __Product__00000015_
	__ProductCompositionShape__00000012_.Product = __Product__00000016_
	__ProductCompositionShape__00000013_.Product = __Product__00000017_
	__ProductCompositionShape__00000014_.Product = __Product__00000018_
	__ProductCompositionShape__00000015_.Product = __Product__00000019_
	__ProductCompositionShape__00000016_.Product = __Product__00000020_
	__ProductCompositionShape__00000017_.Product = __Product__00000021_
	__ProductCompositionShape__00000018_.Product = __Product__00000022_
	__ProductCompositionShape__00000019_.Product = __Product__00000023_
	__ProductCompositionShape__00000020_.Product = __Product__00000024_
	__ProductCompositionShape__00000021_.Product = __Product__00000025_
	__ProductCompositionShape__00000022_.Product = __Product__00000026_
	__ProductCompositionShape__00000023_.Product = __Product__00000027_
	__ProductCompositionShape__00000024_.Product = __Product__00000028_
	__ProductCompositionShape__00000025_.Product = __Product__00000029_
	__ProductShape__00000002_.Product = __Product__00000001_
	__ProductShape__00000003_.Product = __Product__00000002_
	__ProductShape__00000004_.Product = __Product__00000003_
	__ProductShape__00000005_.Product = __Product__00000004_
	__ProductShape__00000006_.Product = __Product__00000005_
	__ProductShape__00000007_.Product = __Product__00000005_
	__ProductShape__00000008_.Product = __Product__00000004_
	__ProductShape__00000009_.Product = __Product__00000006_
	__ProductShape__00000010_.Product = __Product__00000007_
	__ProductShape__00000011_.Product = __Product__00000008_
	__ProductShape__00000012_.Product = __Product__00000005_
	__ProductShape__00000013_.Product = __Product__00000009_
	__ProductShape__00000014_.Product = __Product__00000010_
	__ProductShape__00000015_.Product = __Product__00000011_
	__ProductShape__00000016_.Product = __Product__00000012_
	__ProductShape__00000017_.Product = __Product__00000013_
	__ProductShape__00000018_.Product = __Product__00000002_
	__ProductShape__00000019_.Product = __Product__00000014_
	__ProductShape__00000020_.Product = __Product__00000015_
	__ProductShape__00000021_.Product = __Product__00000016_
	__ProductShape__00000022_.Product = __Product__00000017_
	__ProductShape__00000023_.Product = __Product__00000018_
	__ProductShape__00000024_.Product = __Product__00000019_
	__ProductShape__00000025_.Product = __Product__00000020_
	__ProductShape__00000026_.Product = __Product__00000021_
	__ProductShape__00000027_.Product = __Product__00000022_
	__ProductShape__00000028_.Product = __Product__00000023_
	__ProductShape__00000029_.Product = __Product__00000024_
	__ProductShape__00000030_.Product = __Product__00000025_
	__ProductShape__00000031_.Product = __Product__00000026_
	__ProductShape__00000032_.Product = __Product__00000027_
	__ProductShape__00000033_.Product = __Product__00000028_
	__ProductShape__00000034_.Product = __Product__00000029_
	__Project__00000000_.RootProducts = append(__Project__00000000_.RootProducts, __Product__00000007_)
	__Project__00000000_.RootProducts = append(__Project__00000000_.RootProducts, __Product__00000003_)
	__Project__00000000_.RootTasks = append(__Project__00000000_.RootTasks, __Task__00000001_)
	__Project__00000000_.RootTasks = append(__Project__00000000_.RootTasks, __Task__00000002_)
	__Project__00000000_.RootTasks = append(__Project__00000000_.RootTasks, __Task__00000003_)
	__Project__00000000_.RootResources = append(__Project__00000000_.RootResources, __Resource__00000003_)
	__Project__00000000_.RootResources = append(__Project__00000000_.RootResources, __Resource__00000006_)
	__Project__00000000_.Notes = append(__Project__00000000_.Notes, __Note__00000000_)
	__Project__00000000_.Notes = append(__Project__00000000_.Notes, __Note__00000001_)
	__Project__00000000_.Notes = append(__Project__00000000_.Notes, __Note__00000002_)
	__Project__00000000_.Notes = append(__Project__00000000_.Notes, __Note__00000003_)
	__Project__00000000_.Notes = append(__Project__00000000_.Notes, __Note__00000004_)
	__Project__00000000_.Notes = append(__Project__00000000_.Notes, __Note__00000005_)
	__Project__00000000_.Notes = append(__Project__00000000_.Notes, __Note__00000006_)
	__Project__00000000_.Notes = append(__Project__00000000_.Notes, __Note__00000007_)
	__Project__00000000_.Notes = append(__Project__00000000_.Notes, __Note__00000008_)
	__Project__00000000_.Notes = append(__Project__00000000_.Notes, __Note__00000009_)
	__Project__00000000_.Notes = append(__Project__00000000_.Notes, __Note__00000010_)
	__Project__00000000_.Notes = append(__Project__00000000_.Notes, __Note__00000011_)
	__Project__00000000_.Notes = append(__Project__00000000_.Notes, __Note__00000012_)
	__Project__00000000_.Diagrams = append(__Project__00000000_.Diagrams, __Diagram__00000000_)
	__Project__00000000_.Diagrams = append(__Project__00000000_.Diagrams, __Diagram__00000001_)
	__Project__00000000_.Diagrams = append(__Project__00000000_.Diagrams, __Diagram__00000002_)
	__Project__00000000_.Diagrams = append(__Project__00000000_.Diagrams, __Diagram__00000003_)
	__Project__00000000_.Diagrams = append(__Project__00000000_.Diagrams, __Diagram__00000004_)
	__Project__00000000_.Diagrams = append(__Project__00000000_.Diagrams, __Diagram__00000005_)
	__Project__00000000_.Diagrams = append(__Project__00000000_.Diagrams, __Diagram__00000006_)
	__Resource__00000000_.Tasks = append(__Resource__00000000_.Tasks, __Task__00000001_)
	__Resource__00000003_.SubResources = append(__Resource__00000003_.SubResources, __Resource__00000004_)
	__Resource__00000003_.SubResources = append(__Resource__00000003_.SubResources, __Resource__00000000_)
	__Resource__00000003_.SubResources = append(__Resource__00000003_.SubResources, __Resource__00000008_)
	__Resource__00000004_.SubResources = append(__Resource__00000004_.SubResources, __Resource__00000005_)
	__Resource__00000004_.SubResources = append(__Resource__00000004_.SubResources, __Resource__00000009_)
	__Resource__00000005_.SubResources = append(__Resource__00000005_.SubResources, __Resource__00000001_)
	__Resource__00000005_.SubResources = append(__Resource__00000005_.SubResources, __Resource__00000002_)
	__Resource__00000006_.SubResources = append(__Resource__00000006_.SubResources, __Resource__00000007_)
	__ResourceCompositionShape__00000000_.Resource = __Resource__00000004_
	__ResourceCompositionShape__00000001_.Resource = __Resource__00000005_
	__ResourceCompositionShape__00000002_.Resource = __Resource__00000001_
	__ResourceCompositionShape__00000003_.Resource = __Resource__00000002_
	__ResourceCompositionShape__00000004_.Resource = __Resource__00000000_
	__ResourceCompositionShape__00000005_.Resource = __Resource__00000007_
	__ResourceCompositionShape__00000006_.Resource = __Resource__00000008_
	__ResourceCompositionShape__00000007_.Resource = __Resource__00000009_
	__ResourceShape__00000000_.Resource = __Resource__00000000_
	__ResourceShape__00000001_.Resource = __Resource__00000001_
	__ResourceShape__00000002_.Resource = __Resource__00000002_
	__ResourceShape__00000003_.Resource = __Resource__00000003_
	__ResourceShape__00000004_.Resource = __Resource__00000004_
	__ResourceShape__00000005_.Resource = __Resource__00000005_
	__ResourceShape__00000006_.Resource = __Resource__00000001_
	__ResourceShape__00000007_.Resource = __Resource__00000002_
	__ResourceShape__00000008_.Resource = __Resource__00000000_
	__ResourceShape__00000009_.Resource = __Resource__00000006_
	__ResourceShape__00000010_.Resource = __Resource__00000007_
	__ResourceShape__00000011_.Resource = __Resource__00000008_
	__ResourceShape__00000012_.Resource = __Resource__00000009_
	__ResourceShape__00000013_.Resource = __Resource__00000000_
	__ResourceShape__00000014_.Resource = __Resource__00000000_
	__ResourceTaskShape__00000000_.Resource = __Resource__00000000_
	__ResourceTaskShape__00000000_.Task = __Task__00000001_
	__Root__00000000_.Projects = append(__Root__00000000_.Projects, __Project__00000000_)
	__Task__00000001_.SubTasks = append(__Task__00000001_.SubTasks, __Task__00000004_)
	__Task__00000002_.SubTasks = append(__Task__00000002_.SubTasks, __Task__00000005_)
	__Task__00000002_.SubTasks = append(__Task__00000002_.SubTasks, __Task__00000006_)
	__Task__00000002_.SubTasks = append(__Task__00000002_.SubTasks, __Task__00000000_)
	__Task__00000004_.Inputs = append(__Task__00000004_.Inputs, __Product__00000004_)
	__Task__00000004_.Outputs = append(__Task__00000004_.Outputs, __Product__00000005_)
	__TaskCompositionShape__00000000_.Task = __Task__00000000_
	__TaskCompositionShape__00000001_.Task = __Task__00000000_
	__TaskCompositionShape__00000002_.Task = __Task__00000004_
	__TaskCompositionShape__00000003_.Task = __Task__00000005_
	__TaskCompositionShape__00000004_.Task = __Task__00000006_
	__TaskInputShape__00000000_.Product = __Product__00000004_
	__TaskInputShape__00000000_.Task = __Task__00000004_
	__TaskOutputShape__00000000_.Task = __Task__00000004_
	__TaskOutputShape__00000000_.Product = __Product__00000005_
	__TaskShape__00000000_.Task = __Task__00000000_
	__TaskShape__00000001_.Task = __Task__00000001_
	__TaskShape__00000002_.Task = __Task__00000002_
	__TaskShape__00000003_.Task = __Task__00000001_
	__TaskShape__00000004_.Task = __Task__00000002_
	__TaskShape__00000005_.Task = __Task__00000000_
	__TaskShape__00000006_.Task = __Task__00000003_
	__TaskShape__00000007_.Task = __Task__00000001_
	__TaskShape__00000008_.Task = __Task__00000004_
	__TaskShape__00000009_.Task = __Task__00000004_
	__TaskShape__00000010_.Task = __Task__00000005_
	__TaskShape__00000011_.Task = __Task__00000006_
}
