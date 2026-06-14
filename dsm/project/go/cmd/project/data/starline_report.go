package main

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/project/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var (
	_ time.Time
	_ = slices.Index[[]int, int]
)

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage

	__Diagram__00000001_ := (&models.Diagram{Name: `WBS`}).Stage(stage)
	__Diagram__00000002_ := (&models.Diagram{Name: `PBS`}).Stage(stage)
	__Diagram__00000003_ := (&models.Diagram{Name: `RBS`}).Stage(stage)
	__Diagram__00000004_ := (&models.Diagram{Name: `PIT focus`}).Stage(stage)
	__Diagram__00000005_ := (&models.Diagram{Name: `PIT Report`}).Stage(stage)
	__Diagram__00000006_ := (&models.Diagram{Name: `RCS PBS`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: `Startliner Mishape Report`}).Stage(stage)

	__Note__00000000_ := (&models.Note{Name: `CFT ended in march 2025`}).Stage(stage)
	__Note__00000001_ := (&models.Note{Name: `A thorough review of the STAR report is advised (p14)
`}).Stage(stage)
	__Note__00000002_ := (&models.Note{Name: `NASA utilized a firm fixed price contracting type for CCtCap. 

This was a significant shift from the cost-plus contracting for traditional NASA builds of developmental vehicles. These shifts signified that CCP was not only positioned to be an innovative, first-of-its kind program for NASA, but how it interacted with new and traditional space flight industry providers was setup to be significantly distinct and different.  `}).Stage(stage)
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
	__Note__00000013_ := (&models.Note{Name: `The leading theory for the proximate cause of the CM RCS failure during the CFT is the formulation of carbazic acid, which corrodes stainless steel. The reaction between carbazic acid and stainless steel creates corrosion particulates in the thruster propellant valve, preventing it from opening. `}).Stage(stage)
	__Note__00000014_ := (&models.Note{Name: `Thruster health is monitored via chamber pressure and Fuel/Oxydizer injector temperatures`}).Stage(stage)
	__Note__00000015_ := (&models.Note{Name: `Commanded pulse lengths are neither recorded nor downlinked near real time with telemetry

This complicates efforts to understand what went wrong with SM RCS jets when data suggests that pulse length is correlated to observed soakback temperatures.  `}).Stage(stage)
	__Note__00000016_ := (&models.Note{Name: `Many thruster pulses are less  FDIR in the  The FMCs see chamber pressure over a data bus from the  but the chamber pressure is only sent to the ground with a sample rate for recording every  limiting insight into thruster performance. This does not meet Nyquist Criterion for capturing the chamber pressure signal. As a result, aliasing effects may occur for pulses shorter than  where high-frequency components of the pressure signal are misrepresented or lost, further complicating accurate reconstruction of thruster behavior.  `}).Stage(stage)
	__Note__00000017_ := (&models.Note{Name: ` S2A2 failed at 9050 m (GMT 14:00) and B1A3 failed at 526m (GMT 14:57).`}).Stage(stage)
	__Note__00000018_ := (&models.Note{Name: `At the conclusion of the hotfire, when thruster is demonstrated, a jet is reenabled with jet-fail FDIR inhibited. Through this sequence, four of five jets were recovered. Spacecraft 6DOF control was reestablished in all axes. T

he spacecraft was moded back to auto and the rendezvous/dock was completed.  While the flight rules allow an SM RCS thruster to be used with jet-fail off FDIR inhibited, doing so removes a hazard control.

Whaou.`}).Stage(stage)
	__Note__00000019_ := (&models.Note{Name: ` degraded performance was still observed. This testing demonstrated the issue was not purely transient.

It should have been done before ?!

 In addition to demonstrating transient degradation during high heating, the test program demonstrated an unrecoverable cumulative degradation that was potentially indicative of the observed performance of CFT. `}).Stage(stage)
	__Note__00000020_ := (&models.Note{Name: `B1.3.3.1.1.1 Heat identifies that excessive heating could cause the poppet to extrude, reducing flow. Poppet extrusion was observed in the post-test tear down of the WSTF (White Sands Test Facility) thruster unit.

The initial heating was caused by the internal doghouse temperatures (addressed in C4.1 Thermal) and the follow-on heating induced by B1.3.3.1.4.2 RCS Thermal Soakback. This caused excessive temperatures beyond the capability of the thruster softgoods, causing the Teflon poppet to extrude. The evidence supports this being a contributor to the thruster fail-offs.`}).Stage(stage)
	__Note__00000021_ := (&models.Note{Name: `B1.3.3.1.4.2 RCS Thermal Soakback identifies the possibility that structure of the RCS thruster
retains heat, and the heat is unable to be dissipated from the firings of the thrusters, causing the
propellants to heat rapidly. `}).Stage(stage)
	__Note__00000022_ := (&models.Note{Name: `Intermediate Cause 5: Inadequate Thruster Thermal Models

In May 2019, the PCB accepted the “SM doghouse thermal model (PROP-06) not in critical model
Database” risk, which was generated in 2017. As identified by Passive Thermal Control System
(PTCS), “there was a gap in Boeing thermal modeling. Specifically, the internal componentry within
each of the doghouses isn’t modeled in THERM-16. This area is modeled by Aerojet-Rocketdyne (AR)
per contract with Boeing PROP. The contract only furnishes the results and not the model that
produces the results. a.) AR model needs to undergoes accreditation. There is no NASA insight into
the validity of the results and therefore verification effort is compromised.”

The Therm-11a model did not incorporate OMAC firings, or initial heating that could occur from the
doghouse environment, which is critical given that the aft thrusters have the longest feed line in
comparison to every other thruster.

There was no thermal model available to adequately understand the full impacts of the thermal
environment upon the RCS thrusters, specifically the aft thrusters with the longest feed line tube,
and very close proximity to the OMAC thrusters`}).Stage(stage)
	__Note__00000023_ := (&models.Note{Name: `Intermediate Cause 6: Insufficient Thruster Qualification

SM RCS Thruster Qualification did not cover the flight envelope for temperature and duty cycle (TLYF).`}).Stage(stage)
	__Note__00000024_ := (&models.Note{Name: `Multiple groups, prior to CTF, discussed the lack of mission representative operational duty cycle testing for SM RCS engines.. This concern is one that has been tracked since prior to OFT1 and is specifically reviewed in this intermediate cause because it is widely accepted that operational duty cycle is a primary driver of hardware temperature. NASA Engineering and the CCP Spacecraft Office have consistently identified this qualification gap and recommended additional testing, however it was not incorporated into the pursued plans for risk assessments between flights.

OFT1: The SM RCS Qual Gaps risk was accepted in September 2019 (PCB-19-383), Boeing
RCS/OMAC (Engine) Qual Issues for OFT.

CFT: The qualification gap risk was addressed in April 2023

. The program directive issued at that time resolved to accept
the qualification gap for duty cycles for CFT only. This acceptance was encompassed within the
previously approved elevated 2x5 risk level from PCB-23-053, which was documented during the
IFA closure for OFT-2 related to valve/injector assembly temperature measurements exceeding
qualification limits. These temperature exceedances were categorized as an Unexplained Anomaly
(UA).
Starliner Qualification Risk History
OFT-1 (Orbital Flight Test 1)
The Service Module (SM) RCS Qual Gaps risk was formally accepted in September 2019 (PCB-19-383) regarding Boeing RCS/OMAC (Engine) qualification issues for the flight.

OFT-2 (Orbital Flight Test 2)
During PCB-20-404, addressing the SM RCS Jet Failure, the program decided to use the SM RCS "as is." It was noted that the OFT Mission Elapsed Time (MET) anomaly actually demonstrated the robustness of the RCS thruster design despite long durations of excessive usage. While design modifications were evaluated for the leading theory behind the B2R3 failure, the CCP SDRT report highlighted this as a missed opportunity to further investigate SM RCS thruster failure modes. This is discussed further in later sections regarding organizational factors and anomaly resolution.

CFT (Crew Flight Test)
The qualification gap risk was addressed in April 2023 during PCB-23-100 (CFT: Starliner Prop, OMAC/RCS Hot-Fire Qual Gaps). The program directive resolved to accept the qualification gap for duty cycles for the CFT only. This acceptance was encompassed within a previously approved elevated 2x5 risk level (PCB-23-053), documented during the In-Flight Anomaly (IFA) closure for OFT-2.

The OFT-2 IFA related to valve/injector assembly temperature measurements exceeding qualification limits, which were categorized as an Unexplained Anomaly (UA). The directive acknowledged that these exceedances were attributed to specific operational duty cycles. Consequently, the thermal risk was accepted as a 2x5 risk for potential jet failure. Despite this, engineering teams and the Engineering Review Board (ERB-23-0045-R2) emphasized the importance of testing SM RCS thrusters using flight-like duty cycles.`}).Stage(stage)
	__Note__00000025_ := (&models.Note{Name: `Organizational Factor 1: Insufficient Anomaly Resolution Process Multiple RCS thruster failures occurred on OFT 1 and OFT 2. While fault trees were developed, due diligence to reach direct and root cause was not performed and adequate block closure rationale was often incomplete.

Note : Anomaly Reports (AR) were closed.

Organizational Factor 2: Mischaracterization of Risk in thruster Qual Gaps leads to Flawed Flight Rationale. The thermal risk for Starliner’s SM RCS thrusters was mischaracterized due to reliance on simple models and inadequate qualification testing that failed to replicate mission-representative conditions. Despite clear evidence of thermal soakback and temperature exceedances, risk acceptance proceeded without resolving key environment and duty cycle concerns.`}).Stage(stage)

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
	__NoteProductShape__00000016_ := (&models.NoteProductShape{Name: `The leading theory for the proximate cause of the CM RCS failure during the CFT is the formulation of carbazic acid, which corrodes stainless steel. The reaction between carbazic acid and stainless steel creates corrosion particulates in the thruster propellant valve, preventing it from opening.  to 4.4   Analysis: CM RCS Jet Failure`}).Stage(stage)
	__NoteProductShape__00000017_ := (&models.NoteProductShape{Name: `Thruster health is monitored via chamber pressure and Fuel/Oxydizer injector temperatures to 4.4   Analysis: CM RCS Jet Failure`}).Stage(stage)
	__NoteProductShape__00000018_ := (&models.NoteProductShape{Name: `Commanded pulse lengths are neither recorded nor downlinked near real time with telemetry to 4.5.1 Description of the system`}).Stage(stage)
	__NoteProductShape__00000019_ := (&models.NoteProductShape{Name: `At the conclusion of the hotfire, when thruster is demonstrated, a jet is reenabled with jet-fail FDIR inhibited. Through this sequence, four of five jets were recovered. Spacecraft 6DOF control was reestablished in all axes. The spacecraft was moded back to auto and the rendezvous/dock was completed.  While the flight rules allow an SM RCS thruster to be used with jet-fail off FDIR inhibited, doing so removes a hazard control to 4.5.2   Description of Events and Timeline `}).Stage(stage)
	__NoteProductShape__00000020_ := (&models.NoteProductShape{Name: `Many thruster pulses are less  FDIR in the  The FMCs see chamber pressure over a data bus from the  but the chamber pressure is only sent to the ground with a sample rate for recording every  limiting insight into thruster performance. This does not meet Nyquist Criterion for capturing the chamber pressure signal. As a result, aliasing effects may occur for pulses shorter than  where high-frequency components of the pressure signal are misrepresented or lost, further complicating accurate reconstruction of thruster behavior.   to 4.5.2   Description of Events and Timeline `}).Stage(stage)
	__NoteProductShape__00000021_ := (&models.NoteProductShape{Name: ` degraded performance was still observed. This testing demonstrated the issue was not purely transient. to 4.5.4   Starliner Engine Testing at White Sands Test Facility (WSTF) during CFT`}).Stage(stage)
	__NoteProductShape__00000022_ := (&models.NoteProductShape{Name: `B1.3.3.1.1.1 Heat identifies that excessive heating could cause the poppet to extrude, reducing flow. Poppet extrusion was observed in the post-test tear down of the WSTF (White Sands Test Facility) thruster unit.

The initial heating was caused by the internal doghouse temperatures (addressed in C4.1 Thermal) and the follow-on heating induced by B1.3.3.1.4.2 RCS Thermal Soakback. This caused excessive temperatures beyond the capability of the thruster softgoods, causing the Teflon poppet to extrude. The evidence supports this being a contributor to the thruster fail-offs. to 4.5.5 Fault Tree`}).Stage(stage)
	__NoteProductShape__00000023_ := (&models.NoteProductShape{Name: `B1.3.3.1.4.2 RCS Thermal Soakback identifies the possibility that structure of the RCS thruster
retains heat, and the heat is unable to be dissipated from the firings of the thrusters, causing the
propellants to heat rapidly.  to 4.5.5 Fault Tree`}).Stage(stage)
	__NoteProductShape__00000024_ := (&models.NoteProductShape{Name: `In May 2019, the PCB accepted the “SM doghouse thermal model (PROP-06) not in critical model
Database” risk, which was generated in 2017. As identified by Passive Thermal Control System
(PTCS), “there was a gap in Boeing thermal modeling. Specifically, the internal componentry within
each of the doghouses isn’t modeled in THERM-16. This area is modeled by Aerojet-Rocketdyne (AR)
per contract with Boeing PROP. The contract only furnishes the results and not the model that
produces the results. a.) AR model needs to undergoes accreditation. There is no NASA insight into
the validity of the results and therefore verification effort is compromised.” to Intermediate Cause 5: Inadequate Thruster Thermal Models Inadequate thermal modelling caused insufficient scrutiny for the thermal environment, leading to excessive heating from RCS thermal soakback and integrated heating from OMACs.`}).Stage(stage)
	__NoteProductShape__00000025_ := (&models.NoteProductShape{Name: `Intermediate Cause 6: Insufficient Thruster Qualification

SM RCS Thruster Qualification did not cover the flight envelope for temperature and duty cycle (TLYF). to Intermediate Cause 5: Inadequate Thruster Thermal Models Inadequate thermal modelling caused insufficient scrutiny for the thermal environment, leading to excessive heating from RCS thermal soakback and integrated heating from OMACs.`}).Stage(stage)
	__NoteProductShape__00000026_ := (&models.NoteProductShape{Name: `Multiple groups, prior to CTF, discussed the lack of mission representative operational duty cycle testing for SM RCS engines.. This concern is one that has been tracked since prior to OFT1 and is specifically reviewed in this intermediate cause because it is widely accepted that operational duty cycle is a primary driver of hardware temperature. NASA Engineering and the CCP Spacecraft Office have consistently identified this qualification gap and recommended additional testing, however it was not incorporated into the pursued plans for risk assessments between flights. to Intermediate Cause 5: Inadequate Thruster Thermal Models Inadequate thermal modelling caused insufficient scrutiny for the thermal environment, leading to excessive heating from RCS thermal soakback and integrated heating from OMACs.`}).Stage(stage)

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
	__NoteShape__00000013_ := (&models.NoteShape{Name: `-PIT Report`}).Stage(stage)
	__NoteShape__00000014_ := (&models.NoteShape{Name: `-PIT Report`}).Stage(stage)
	__NoteShape__00000015_ := (&models.NoteShape{Name: `-PIT Report`}).Stage(stage)
	__NoteShape__00000016_ := (&models.NoteShape{Name: `-PIT Report`}).Stage(stage)
	__NoteShape__00000017_ := (&models.NoteShape{Name: `-WBS`}).Stage(stage)
	__NoteShape__00000018_ := (&models.NoteShape{Name: `-PIT Report`}).Stage(stage)
	__NoteShape__00000019_ := (&models.NoteShape{Name: `-PIT Report`}).Stage(stage)
	__NoteShape__00000020_ := (&models.NoteShape{Name: `-PIT Report`}).Stage(stage)
	__NoteShape__00000021_ := (&models.NoteShape{Name: `-PIT Report`}).Stage(stage)
	__NoteShape__00000022_ := (&models.NoteShape{Name: `-PIT Report`}).Stage(stage)
	__NoteShape__00000023_ := (&models.NoteShape{Name: `-PIT Report`}).Stage(stage)
	__NoteShape__00000024_ := (&models.NoteShape{Name: `-PIT Report`}).Stage(stage)
	__NoteShape__00000025_ := (&models.NoteShape{Name: `-PIT Report`}).Stage(stage)

	__NoteTaskShape__00000001_ := (&models.NoteTaskShape{Name: ` S2A2 failed at 9050 m (GMT 14:00) and B1A3 failed at 526m (GMT 14:57). to ISS Approach`}).Stage(stage)

	__Product__00000001_ := (&models.Product{Name: `Dragon`}).Stage(stage)
	__Product__00000002_ := (&models.Product{Name: `CST-100 Starliner`}).Stage(stage)
	__Product__00000003_ := (&models.Product{Name: `Reports`}).Stage(stage)
	__Product__00000004_ := (&models.Product{Name: ` Starliner Tests and Anomalies Review (STAR) Investigation Report`}).Stage(stage)
	__Product__00000005_ := (&models.Product{Name: `Program Investigation Team (PIT) Reports`}).Stage(stage)
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
	__Product__00000030_ := (&models.Product{Name: `4.1   Objectives and Approach `}).Stage(stage)
	__Product__00000031_ := (&models.Product{Name: `4.2 Definitions`}).Stage(stage)
	__Product__00000032_ := (&models.Product{Name: `4.3 Fault Tree`}).Stage(stage)
	__Product__00000033_ := (&models.Product{Name: `Guidance, Navigation, and Control (GNC)`}).Stage(stage)
	__Product__00000034_ := (&models.Product{Name: `4.4   Analysis: CM RCS Jet Failure`}).Stage(stage)
	__Product__00000035_ := (&models.Product{Name: `4.5   Analysis: Loss of 6DOF Control - SM RCS Jet Failures`}).Stage(stage)
	__Product__00000036_ := (&models.Product{Name: `4.5.1 Description of the system`}).Stage(stage)
	__Product__00000037_ := (&models.Product{Name: ``}).Stage(stage)
	__Product__00000038_ := (&models.Product{Name: `Inertial Measurement Unit (IMU) `}).Stage(stage)
	__Product__00000039_ := (&models.Product{Name: `4.5.2   Description of Events and Timeline `}).Stage(stage)
	__Product__00000040_ := (&models.Product{Name: `B1A3 Thruster failure`}).Stage(stage)
	__Product__00000041_ := (&models.Product{Name: `4.5.4   Starliner Engine Testing at White Sands Test Facility (WSTF) during CFT`}).Stage(stage)
	__Product__00000042_ := (&models.Product{Name: `4.5.5 Fault Tree`}).Stage(stage)
	__Product__00000043_ := (&models.Product{Name: `4.5.6 Most Probable Proximate Cause`}).Stage(stage)
	__Product__00000044_ := (&models.Product{Name: `Intermediate Cause 5: Inadequate Thruster Thermal Models Inadequate thermal modelling caused insufficient scrutiny for the thermal environment, leading to excessive heating from RCS thermal soakback and integrated heating from OMACs.`}).Stage(stage)
	__Product__00000045_ := (&models.Product{Name: `4.6 Analysis: Helium Leak`}).Stage(stage)
	__Product__00000046_ := (&models.Product{Name: `4.6.1 Description of the system`}).Stage(stage)

	__ProductCompositionShape__00000000_ := (&models.ProductCompositionShape{Name: `Reports to `}).Stage(stage)
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
	__ProductCompositionShape__00000016_ := (&models.ProductCompositionShape{Name: `12 CM RCS Thrusters to `}).Stage(stage)
	__ProductCompositionShape__00000018_ := (&models.ProductCompositionShape{Name: `Reports to Boeing’s Enterprise Root Cause/Corrective Actions (eRCCA)`}).Stage(stage)
	__ProductCompositionShape__00000019_ := (&models.ProductCompositionShape{Name: `Chap3 Commercial Crew Program (CCP) Background to `}).Stage(stage)
	__ProductCompositionShape__00000020_ := (&models.ProductCompositionShape{Name: `Chap3 Commercial Crew Program (CCP) Background to `}).Stage(stage)
	__ProductCompositionShape__00000021_ := (&models.ProductCompositionShape{Name: `CST-100 Starliner to `}).Stage(stage)
	__ProductCompositionShape__00000022_ := (&models.ProductCompositionShape{Name: `CST-100 Starliner to `}).Stage(stage)
	__ProductCompositionShape__00000023_ := (&models.ProductCompositionShape{Name: `Chap3 Commercial Crew Program (CCP) Background to `}).Stage(stage)
	__ProductCompositionShape__00000024_ := (&models.ProductCompositionShape{Name: `12 CM RCS Thrusters (1200 lbf) to `}).Stage(stage)
	__ProductCompositionShape__00000025_ := (&models.ProductCompositionShape{Name: `Chap3 Commercial Crew Program (CCP) Background to `}).Stage(stage)
	__ProductCompositionShape__00000026_ := (&models.ProductCompositionShape{Name: `4     Technical Root Cause Analysis (RCA) and Findings to `}).Stage(stage)
	__ProductCompositionShape__00000027_ := (&models.ProductCompositionShape{Name: `4     Technical Root Cause Analysis (RCA) and Findings to `}).Stage(stage)
	__ProductCompositionShape__00000028_ := (&models.ProductCompositionShape{Name: `4     Technical Root Cause Analysis (RCA) and Findings to `}).Stage(stage)
	__ProductCompositionShape__00000029_ := (&models.ProductCompositionShape{Name: `CST-100 Starliner to `}).Stage(stage)
	__ProductCompositionShape__00000030_ := (&models.ProductCompositionShape{Name: `4     Technical Root Cause Analysis (RCA) and Findings to `}).Stage(stage)
	__ProductCompositionShape__00000031_ := (&models.ProductCompositionShape{Name: `4     Technical Root Cause Analysis (RCA) and Findings to `}).Stage(stage)
	__ProductCompositionShape__00000032_ := (&models.ProductCompositionShape{Name: `4.5   Analysis: Loss of 6DOF Control - SM RCS Jet Failures to `}).Stage(stage)
	__ProductCompositionShape__00000033_ := (&models.ProductCompositionShape{Name: `4     Technical Root Cause Analysis (RCA) and Findings to `}).Stage(stage)
	__ProductCompositionShape__00000034_ := (&models.ProductCompositionShape{Name: `Guidance, Navigation, and Control (GNC) to `}).Stage(stage)
	__ProductCompositionShape__00000035_ := (&models.ProductCompositionShape{Name: `CST-100 Starliner to Orbital Maneuvering and Attitude Control (OMAC), 30 000 lbf`}).Stage(stage)
	__ProductCompositionShape__00000036_ := (&models.ProductCompositionShape{Name: `4.5   Analysis: Loss of 6DOF Control - SM RCS Jet Failures to `}).Stage(stage)
	__ProductCompositionShape__00000037_ := (&models.ProductCompositionShape{Name: `4.5   Analysis: Loss of 6DOF Control - SM RCS Jet Failures to `}).Stage(stage)
	__ProductCompositionShape__00000038_ := (&models.ProductCompositionShape{Name: `4.5   Analysis: Loss of 6DOF Control - SM RCS Jet Failures to `}).Stage(stage)
	__ProductCompositionShape__00000039_ := (&models.ProductCompositionShape{Name: `4.5   Analysis: Loss of 6DOF Control - SM RCS Jet Failures to `}).Stage(stage)
	__ProductCompositionShape__00000040_ := (&models.ProductCompositionShape{Name: `4.5   Analysis: Loss of 6DOF Control - SM RCS Jet Failures to `}).Stage(stage)
	__ProductCompositionShape__00000041_ := (&models.ProductCompositionShape{Name: `4.5.6 Most Probable Proximate Cause to `}).Stage(stage)
	__ProductCompositionShape__00000042_ := (&models.ProductCompositionShape{Name: `4.5   Analysis: Loss of 6DOF Control - SM RCS Jet Failures to `}).Stage(stage)
	__ProductCompositionShape__00000043_ := (&models.ProductCompositionShape{Name: `4.6 Analysis: Helium Leak to `}).Stage(stage)
	__ProductCompositionShape__00000044_ := (&models.ProductCompositionShape{Name: `Reports to Program Investigation Team (PIT) Reports`}).Stage(stage)

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
	__ProductShape__00000027_ := (&models.ProductShape{Name: `Boeing’s Enterprise Root Cause/Corrective Actions (eRCCA)-PBS`}).Stage(stage)
	__ProductShape__00000028_ := (&models.ProductShape{Name: `-PIT Report`}).Stage(stage)
	__ProductShape__00000029_ := (&models.ProductShape{Name: `-PIT Report`}).Stage(stage)
	__ProductShape__00000030_ := (&models.ProductShape{Name: `-RCS PBS`}).Stage(stage)
	__ProductShape__00000031_ := (&models.ProductShape{Name: `-RCS PBS`}).Stage(stage)
	__ProductShape__00000032_ := (&models.ProductShape{Name: `-PIT Report`}).Stage(stage)
	__ProductShape__00000033_ := (&models.ProductShape{Name: `-RCS PBS`}).Stage(stage)
	__ProductShape__00000034_ := (&models.ProductShape{Name: `-PIT Report`}).Stage(stage)
	__ProductShape__00000035_ := (&models.ProductShape{Name: `-PIT Report`}).Stage(stage)
	__ProductShape__00000036_ := (&models.ProductShape{Name: `-PIT Report`}).Stage(stage)
	__ProductShape__00000037_ := (&models.ProductShape{Name: `-PIT Report`}).Stage(stage)
	__ProductShape__00000038_ := (&models.ProductShape{Name: `-RCS PBS`}).Stage(stage)
	__ProductShape__00000039_ := (&models.ProductShape{Name: `-PIT Report`}).Stage(stage)
	__ProductShape__00000040_ := (&models.ProductShape{Name: `-PIT Report`}).Stage(stage)
	__ProductShape__00000041_ := (&models.ProductShape{Name: `-PIT Report`}).Stage(stage)
	__ProductShape__00000042_ := (&models.ProductShape{Name: `-PIT Report`}).Stage(stage)
	__ProductShape__00000043_ := (&models.ProductShape{Name: `-RCS PBS`}).Stage(stage)
	__ProductShape__00000044_ := (&models.ProductShape{Name: `Orbital Maneuvering and Attitude Control (OMAC), 30 000 lbf-RCS PBS`}).Stage(stage)
	__ProductShape__00000045_ := (&models.ProductShape{Name: `-PIT Report`}).Stage(stage)
	__ProductShape__00000046_ := (&models.ProductShape{Name: `-PIT Report`}).Stage(stage)
	__ProductShape__00000047_ := (&models.ProductShape{Name: `-PIT Report`}).Stage(stage)
	__ProductShape__00000048_ := (&models.ProductShape{Name: `-PIT Report`}).Stage(stage)
	__ProductShape__00000049_ := (&models.ProductShape{Name: `-PIT Report`}).Stage(stage)
	__ProductShape__00000050_ := (&models.ProductShape{Name: `-PIT Report`}).Stage(stage)
	__ProductShape__00000051_ := (&models.ProductShape{Name: `-PIT Report`}).Stage(stage)
	__ProductShape__00000052_ := (&models.ProductShape{Name: `-PIT Report`}).Stage(stage)

	__Resource__00000000_ := (&models.Resource{Name: `Program Investigation Team (PIT)`}).Stage(stage)
	__Resource__00000001_ := (&models.Resource{Name: `Barry "Butch" Wilmore`}).Stage(stage)
	__Resource__00000002_ := (&models.Resource{Name: `Sunita "Suni" Williams`}).Stage(stage)
	__Resource__00000003_ := (&models.Resource{Name: `NASA`}).Stage(stage)
	__Resource__00000004_ := (&models.Resource{Name: `Crew Commercial Program (CPP2)`}).Stage(stage)
	__Resource__00000005_ := (&models.Resource{Name: `Crews`}).Stage(stage)
	__Resource__00000006_ := (&models.Resource{Name: `Boeing`}).Stage(stage)
	__Resource__00000007_ := (&models.Resource{Name: ``}).Stage(stage)
	__Resource__00000009_ := (&models.Resource{Name: ` Starliner Tests and Anomalies Review (STAR) Investigation Team`}).Stage(stage)
	__Resource__00000010_ := (&models.Resource{Name: `Flight Control Team (FCT)`}).Stage(stage)

	__ResourceCompositionShape__00000000_ := (&models.ResourceCompositionShape{Name: `NASA to `}).Stage(stage)
	__ResourceCompositionShape__00000001_ := (&models.ResourceCompositionShape{Name: `Crew Commercial Program (CPP) to `}).Stage(stage)
	__ResourceCompositionShape__00000002_ := (&models.ResourceCompositionShape{Name: `Crews to Barry "Butch" Wilmore`}).Stage(stage)
	__ResourceCompositionShape__00000003_ := (&models.ResourceCompositionShape{Name: `Crews to Sunita "Suni" Williams`}).Stage(stage)
	__ResourceCompositionShape__00000004_ := (&models.ResourceCompositionShape{Name: `NASA to PITProgram Investigation Team (PIT)`}).Stage(stage)
	__ResourceCompositionShape__00000005_ := (&models.ResourceCompositionShape{Name: `Boeing to `}).Stage(stage)
	__ResourceCompositionShape__00000007_ := (&models.ResourceCompositionShape{Name: `Crew Commercial Program (CPP) to `}).Stage(stage)
	__ResourceCompositionShape__00000008_ := (&models.ResourceCompositionShape{Name: `Crew Commercial Program (CPP) to `}).Stage(stage)

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
	__ResourceShape__00000012_ := (&models.ResourceShape{Name: `-RBS`}).Stage(stage)
	__ResourceShape__00000013_ := (&models.ResourceShape{Name: `Program Investigation Team (PIT)-PBS`}).Stage(stage)
	__ResourceShape__00000014_ := (&models.ResourceShape{Name: `Program Investigation Team (PIT)-PIT focus`}).Stage(stage)
	__ResourceShape__00000015_ := (&models.ResourceShape{Name: `-RBS`}).Stage(stage)

	__ResourceTaskShape__00000000_ := (&models.ResourceTaskShape{Name: `PITProgram Investigation Team (PIT) to Mishap investigation`}).Stage(stage)

	__Task__00000000_ := (&models.Task{Name: `Starliner Crewed Flight Test (CFT)`}).Stage(stage)
	__Task__00000001_ := (&models.Task{Name: `Mishap investigations`}).Stage(stage)
	__Task__00000002_ := (&models.Task{Name: `Commercial Crew Program (CCP),`}).Stage(stage)
	__Task__00000003_ := (&models.Task{Name: ` Commercial ReSupply (CRS) `}).Stage(stage)
	__Task__00000004_ := (&models.Task{Name: `Program Investigation Team (PIT) Report`}).Stage(stage)
	__Task__00000005_ := (&models.Task{Name: `Orbital Flight Test-1 (OFT-1)`}).Stage(stage)
	__Task__00000006_ := (&models.Task{Name: `Orbital Flight Test-2 (OFT-2)`}).Stage(stage)
	__Task__00000007_ := (&models.Task{Name: `ISS Approach`}).Stage(stage)

	__TaskCompositionShape__00000001_ := (&models.TaskCompositionShape{Name: `Commercial Crew Program (CCP), to Starliner Crewed Flight Test (CFT)`}).Stage(stage)
	__TaskCompositionShape__00000002_ := (&models.TaskCompositionShape{Name: `Mishap investigations to `}).Stage(stage)
	__TaskCompositionShape__00000003_ := (&models.TaskCompositionShape{Name: `Commercial Crew Program (CCP), to `}).Stage(stage)
	__TaskCompositionShape__00000004_ := (&models.TaskCompositionShape{Name: `Commercial Crew Program (CCP), to `}).Stage(stage)
	__TaskCompositionShape__00000005_ := (&models.TaskCompositionShape{Name: `Starliner Crewed Flight Test (CFT) to `}).Stage(stage)

	__TaskInputShape__00000000_ := (&models.TaskInputShape{Name: `Program Investigation Team (PIT) Report to  Starliner Tests and Anomalies Review (STAR) Investigation Report`}).Stage(stage)

	__TaskOutputShape__00000000_ := (&models.TaskOutputShape{Name: `Program Investigation Team (PIT) Report to Program Investigation Team (PIT) Report`}).Stage(stage)

	__TaskShape__00000003_ := (&models.TaskShape{Name: `Mishap investigation-WBS`}).Stage(stage)
	__TaskShape__00000004_ := (&models.TaskShape{Name: `Commercial Crew Program (CCP),-WBS`}).Stage(stage)
	__TaskShape__00000005_ := (&models.TaskShape{Name: `Starliner Crewed Flight Test (CFT)-WBS`}).Stage(stage)
	__TaskShape__00000006_ := (&models.TaskShape{Name: `-WBS`}).Stage(stage)
	__TaskShape__00000007_ := (&models.TaskShape{Name: `Mishap investigation-PBS`}).Stage(stage)
	__TaskShape__00000008_ := (&models.TaskShape{Name: `-WBS`}).Stage(stage)
	__TaskShape__00000009_ := (&models.TaskShape{Name: `Program Investigation Team (PIT) Report-PIT focus`}).Stage(stage)
	__TaskShape__00000010_ := (&models.TaskShape{Name: `-WBS`}).Stage(stage)
	__TaskShape__00000011_ := (&models.TaskShape{Name: `-WBS`}).Stage(stage)
	__TaskShape__00000012_ := (&models.TaskShape{Name: `-WBS`}).Stage(stage)

	// insertion point for initialization of values

	__Diagram__00000001_.Name = `WBS`
	__Diagram__00000001_.ComputedPrefix = `1`
	__Diagram__00000001_.IsExpanded = false
	__Diagram__00000001_.IsChecked = true
	__Diagram__00000001_.IsEditable_ = true
	__Diagram__00000001_.IsShowPrefix = true
	__Diagram__00000001_.DefaultBoxWidth = 250.000000
	__Diagram__00000001_.DefaultBoxHeigth = 70.000000
	__Diagram__00000001_.Width = 995.617766
	__Diagram__00000001_.Height = 856.674065
	__Diagram__00000001_.IsPBSNodeExpanded = true
	__Diagram__00000001_.IsWBSNodeExpanded = false
	__Diagram__00000001_.IsTaskGroupsNodeExpanded = false
	__Diagram__00000001_.IsMilestonesNodeExpanded = false
	__Diagram__00000001_.DateFormat = ``
	__Diagram__00000001_.IsNotesNodeExpanded = false
	__Diagram__00000001_.IsResourcesNodeExpanded = true
	__Diagram__00000001_.IsTimeDiagram = false
	__Diagram__00000001_.ComputedStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000001_.ComputedEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000001_.ComputedDuration = 0
	__Diagram__00000001_.UseManualStartAndEndDates = false
	__Diagram__00000001_.ManualStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000001_.ManualEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000001_.TimeStep = 0
	__Diagram__00000001_.TimeStepScale = ""
	__Diagram__00000001_.LaneHeight = 0.000000
	__Diagram__00000001_.RatioBarToLaneHeight = 0.000000
	__Diagram__00000001_.YTopMargin = 0.000000
	__Diagram__00000001_.XLeftText = 0.000000
	__Diagram__00000001_.TextHeight = 0.000000
	__Diagram__00000001_.XLeftLanes = 0.000000
	__Diagram__00000001_.XRightMargin = 0.000000
	__Diagram__00000001_.ArrowLengthToTheRightOfStartBar = 0.000000
	__Diagram__00000001_.ArrowTipLenght = 0.000000
	__Diagram__00000001_.TimeLine_Color = ``
	__Diagram__00000001_.TimeLine_FillOpacity = 0.000000
	__Diagram__00000001_.TimeLine_Stroke = ``
	__Diagram__00000001_.TimeLine_StrokeWidth = 0.000000
	__Diagram__00000001_.DrawVerticalTimeLines = false
	__Diagram__00000001_.Group_Stroke = ``
	__Diagram__00000001_.Group_StrokeWidth = 0.000000
	__Diagram__00000001_.Group_StrokeDashArray = ``
	__Diagram__00000001_.DateYOffset = 0.000000
	__Diagram__00000001_.AlignOnStartEndOnYearStart = false

	__Diagram__00000002_.Name = `PBS`
	__Diagram__00000002_.ComputedPrefix = `2`
	__Diagram__00000002_.IsExpanded = true
	__Diagram__00000002_.IsChecked = false
	__Diagram__00000002_.IsEditable_ = true
	__Diagram__00000002_.IsShowPrefix = true
	__Diagram__00000002_.DefaultBoxWidth = 250.000000
	__Diagram__00000002_.DefaultBoxHeigth = 70.000000
	__Diagram__00000002_.Width = 1475.000000
	__Diagram__00000002_.Height = 1005.999985
	__Diagram__00000002_.IsPBSNodeExpanded = true
	__Diagram__00000002_.IsWBSNodeExpanded = false
	__Diagram__00000002_.IsTaskGroupsNodeExpanded = false
	__Diagram__00000002_.IsMilestonesNodeExpanded = false
	__Diagram__00000002_.DateFormat = ``
	__Diagram__00000002_.IsNotesNodeExpanded = false
	__Diagram__00000002_.IsResourcesNodeExpanded = false
	__Diagram__00000002_.IsTimeDiagram = false
	__Diagram__00000002_.ComputedStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000002_.ComputedEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000002_.ComputedDuration = 0
	__Diagram__00000002_.UseManualStartAndEndDates = false
	__Diagram__00000002_.ManualStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000002_.ManualEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000002_.TimeStep = 0
	__Diagram__00000002_.TimeStepScale = ""
	__Diagram__00000002_.LaneHeight = 0.000000
	__Diagram__00000002_.RatioBarToLaneHeight = 0.000000
	__Diagram__00000002_.YTopMargin = 0.000000
	__Diagram__00000002_.XLeftText = 0.000000
	__Diagram__00000002_.TextHeight = 0.000000
	__Diagram__00000002_.XLeftLanes = 0.000000
	__Diagram__00000002_.XRightMargin = 0.000000
	__Diagram__00000002_.ArrowLengthToTheRightOfStartBar = 0.000000
	__Diagram__00000002_.ArrowTipLenght = 0.000000
	__Diagram__00000002_.TimeLine_Color = ``
	__Diagram__00000002_.TimeLine_FillOpacity = 0.000000
	__Diagram__00000002_.TimeLine_Stroke = ``
	__Diagram__00000002_.TimeLine_StrokeWidth = 0.000000
	__Diagram__00000002_.DrawVerticalTimeLines = false
	__Diagram__00000002_.Group_Stroke = ``
	__Diagram__00000002_.Group_StrokeWidth = 0.000000
	__Diagram__00000002_.Group_StrokeDashArray = ``
	__Diagram__00000002_.DateYOffset = 0.000000
	__Diagram__00000002_.AlignOnStartEndOnYearStart = false

	__Diagram__00000003_.Name = `RBS`
	__Diagram__00000003_.ComputedPrefix = `3`
	__Diagram__00000003_.IsExpanded = true
	__Diagram__00000003_.IsChecked = false
	__Diagram__00000003_.IsEditable_ = true
	__Diagram__00000003_.IsShowPrefix = true
	__Diagram__00000003_.DefaultBoxWidth = 250.000000
	__Diagram__00000003_.DefaultBoxHeigth = 70.000000
	__Diagram__00000003_.Width = 1233.737820
	__Diagram__00000003_.Height = 1199.201740
	__Diagram__00000003_.IsPBSNodeExpanded = true
	__Diagram__00000003_.IsWBSNodeExpanded = false
	__Diagram__00000003_.IsTaskGroupsNodeExpanded = false
	__Diagram__00000003_.IsMilestonesNodeExpanded = false
	__Diagram__00000003_.DateFormat = ``
	__Diagram__00000003_.IsNotesNodeExpanded = false
	__Diagram__00000003_.IsResourcesNodeExpanded = false
	__Diagram__00000003_.IsTimeDiagram = false
	__Diagram__00000003_.ComputedStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000003_.ComputedEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000003_.ComputedDuration = 0
	__Diagram__00000003_.UseManualStartAndEndDates = false
	__Diagram__00000003_.ManualStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000003_.ManualEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000003_.TimeStep = 0
	__Diagram__00000003_.TimeStepScale = ""
	__Diagram__00000003_.LaneHeight = 0.000000
	__Diagram__00000003_.RatioBarToLaneHeight = 0.000000
	__Diagram__00000003_.YTopMargin = 0.000000
	__Diagram__00000003_.XLeftText = 0.000000
	__Diagram__00000003_.TextHeight = 0.000000
	__Diagram__00000003_.XLeftLanes = 0.000000
	__Diagram__00000003_.XRightMargin = 0.000000
	__Diagram__00000003_.ArrowLengthToTheRightOfStartBar = 0.000000
	__Diagram__00000003_.ArrowTipLenght = 0.000000
	__Diagram__00000003_.TimeLine_Color = ``
	__Diagram__00000003_.TimeLine_FillOpacity = 0.000000
	__Diagram__00000003_.TimeLine_Stroke = ``
	__Diagram__00000003_.TimeLine_StrokeWidth = 0.000000
	__Diagram__00000003_.DrawVerticalTimeLines = false
	__Diagram__00000003_.Group_Stroke = ``
	__Diagram__00000003_.Group_StrokeWidth = 0.000000
	__Diagram__00000003_.Group_StrokeDashArray = ``
	__Diagram__00000003_.DateYOffset = 0.000000
	__Diagram__00000003_.AlignOnStartEndOnYearStart = false

	__Diagram__00000004_.Name = `PIT focus`
	__Diagram__00000004_.ComputedPrefix = `4`
	__Diagram__00000004_.IsExpanded = false
	__Diagram__00000004_.IsChecked = false
	__Diagram__00000004_.IsEditable_ = true
	__Diagram__00000004_.IsShowPrefix = true
	__Diagram__00000004_.DefaultBoxWidth = 250.000000
	__Diagram__00000004_.DefaultBoxHeigth = 70.000000
	__Diagram__00000004_.Width = 1262.017987
	__Diagram__00000004_.Height = 458.658205
	__Diagram__00000004_.IsPBSNodeExpanded = true
	__Diagram__00000004_.IsWBSNodeExpanded = false
	__Diagram__00000004_.IsTaskGroupsNodeExpanded = false
	__Diagram__00000004_.IsMilestonesNodeExpanded = false
	__Diagram__00000004_.DateFormat = ``
	__Diagram__00000004_.IsNotesNodeExpanded = true
	__Diagram__00000004_.IsResourcesNodeExpanded = true
	__Diagram__00000004_.IsTimeDiagram = false
	__Diagram__00000004_.ComputedStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000004_.ComputedEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000004_.ComputedDuration = 0
	__Diagram__00000004_.UseManualStartAndEndDates = false
	__Diagram__00000004_.ManualStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000004_.ManualEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000004_.TimeStep = 0
	__Diagram__00000004_.TimeStepScale = ""
	__Diagram__00000004_.LaneHeight = 0.000000
	__Diagram__00000004_.RatioBarToLaneHeight = 0.000000
	__Diagram__00000004_.YTopMargin = 0.000000
	__Diagram__00000004_.XLeftText = 0.000000
	__Diagram__00000004_.TextHeight = 0.000000
	__Diagram__00000004_.XLeftLanes = 0.000000
	__Diagram__00000004_.XRightMargin = 0.000000
	__Diagram__00000004_.ArrowLengthToTheRightOfStartBar = 0.000000
	__Diagram__00000004_.ArrowTipLenght = 0.000000
	__Diagram__00000004_.TimeLine_Color = ``
	__Diagram__00000004_.TimeLine_FillOpacity = 0.000000
	__Diagram__00000004_.TimeLine_Stroke = ``
	__Diagram__00000004_.TimeLine_StrokeWidth = 0.000000
	__Diagram__00000004_.DrawVerticalTimeLines = false
	__Diagram__00000004_.Group_Stroke = ``
	__Diagram__00000004_.Group_StrokeWidth = 0.000000
	__Diagram__00000004_.Group_StrokeDashArray = ``
	__Diagram__00000004_.DateYOffset = 0.000000
	__Diagram__00000004_.AlignOnStartEndOnYearStart = false

	__Diagram__00000005_.Name = `PIT Report`
	__Diagram__00000005_.ComputedPrefix = `5`
	__Diagram__00000005_.IsExpanded = false
	__Diagram__00000005_.IsChecked = false
	__Diagram__00000005_.IsEditable_ = true
	__Diagram__00000005_.IsShowPrefix = true
	__Diagram__00000005_.DefaultBoxWidth = 250.000000
	__Diagram__00000005_.DefaultBoxHeigth = 70.000000
	__Diagram__00000005_.Width = 1739.887483
	__Diagram__00000005_.Height = 3348.225340
	__Diagram__00000005_.IsPBSNodeExpanded = true
	__Diagram__00000005_.IsWBSNodeExpanded = false
	__Diagram__00000005_.IsTaskGroupsNodeExpanded = false
	__Diagram__00000005_.IsMilestonesNodeExpanded = false
	__Diagram__00000005_.DateFormat = ``
	__Diagram__00000005_.IsNotesNodeExpanded = true
	__Diagram__00000005_.IsResourcesNodeExpanded = false
	__Diagram__00000005_.IsTimeDiagram = false
	__Diagram__00000005_.ComputedStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000005_.ComputedEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000005_.ComputedDuration = 0
	__Diagram__00000005_.UseManualStartAndEndDates = false
	__Diagram__00000005_.ManualStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000005_.ManualEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000005_.TimeStep = 0
	__Diagram__00000005_.TimeStepScale = ""
	__Diagram__00000005_.LaneHeight = 0.000000
	__Diagram__00000005_.RatioBarToLaneHeight = 0.000000
	__Diagram__00000005_.YTopMargin = 0.000000
	__Diagram__00000005_.XLeftText = 0.000000
	__Diagram__00000005_.TextHeight = 0.000000
	__Diagram__00000005_.XLeftLanes = 0.000000
	__Diagram__00000005_.XRightMargin = 0.000000
	__Diagram__00000005_.ArrowLengthToTheRightOfStartBar = 0.000000
	__Diagram__00000005_.ArrowTipLenght = 0.000000
	__Diagram__00000005_.TimeLine_Color = ``
	__Diagram__00000005_.TimeLine_FillOpacity = 0.000000
	__Diagram__00000005_.TimeLine_Stroke = ``
	__Diagram__00000005_.TimeLine_StrokeWidth = 0.000000
	__Diagram__00000005_.DrawVerticalTimeLines = false
	__Diagram__00000005_.Group_Stroke = ``
	__Diagram__00000005_.Group_StrokeWidth = 0.000000
	__Diagram__00000005_.Group_StrokeDashArray = ``
	__Diagram__00000005_.DateYOffset = 0.000000
	__Diagram__00000005_.AlignOnStartEndOnYearStart = false

	__Diagram__00000006_.Name = `RCS PBS`
	__Diagram__00000006_.ComputedPrefix = `6`
	__Diagram__00000006_.IsExpanded = false
	__Diagram__00000006_.IsChecked = false
	__Diagram__00000006_.IsEditable_ = true
	__Diagram__00000006_.IsShowPrefix = false
	__Diagram__00000006_.DefaultBoxWidth = 250.000000
	__Diagram__00000006_.DefaultBoxHeigth = 70.000000
	__Diagram__00000006_.Width = 1569.153628
	__Diagram__00000006_.Height = 1277.529266
	__Diagram__00000006_.IsPBSNodeExpanded = true
	__Diagram__00000006_.IsWBSNodeExpanded = false
	__Diagram__00000006_.IsTaskGroupsNodeExpanded = false
	__Diagram__00000006_.IsMilestonesNodeExpanded = false
	__Diagram__00000006_.DateFormat = ``
	__Diagram__00000006_.IsNotesNodeExpanded = true
	__Diagram__00000006_.IsResourcesNodeExpanded = false
	__Diagram__00000006_.IsTimeDiagram = false
	__Diagram__00000006_.ComputedStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000006_.ComputedEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000006_.ComputedDuration = 0
	__Diagram__00000006_.UseManualStartAndEndDates = false
	__Diagram__00000006_.ManualStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000006_.ManualEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000006_.TimeStep = 0
	__Diagram__00000006_.TimeStepScale = ""
	__Diagram__00000006_.LaneHeight = 0.000000
	__Diagram__00000006_.RatioBarToLaneHeight = 0.000000
	__Diagram__00000006_.YTopMargin = 0.000000
	__Diagram__00000006_.XLeftText = 0.000000
	__Diagram__00000006_.TextHeight = 0.000000
	__Diagram__00000006_.XLeftLanes = 0.000000
	__Diagram__00000006_.XRightMargin = 0.000000
	__Diagram__00000006_.ArrowLengthToTheRightOfStartBar = 0.000000
	__Diagram__00000006_.ArrowTipLenght = 0.000000
	__Diagram__00000006_.TimeLine_Color = ``
	__Diagram__00000006_.TimeLine_FillOpacity = 0.000000
	__Diagram__00000006_.TimeLine_Stroke = ``
	__Diagram__00000006_.TimeLine_StrokeWidth = 0.000000
	__Diagram__00000006_.DrawVerticalTimeLines = false
	__Diagram__00000006_.Group_Stroke = ``
	__Diagram__00000006_.Group_StrokeWidth = 0.000000
	__Diagram__00000006_.Group_StrokeDashArray = ``
	__Diagram__00000006_.DateYOffset = 0.000000
	__Diagram__00000006_.AlignOnStartEndOnYearStart = false

	__Library__00000000_.Name = `Startliner Mishape Report`
	__Library__00000000_.NbPixPerCharacter = 8.000000
	__Library__00000000_.LogoSVGFile = `<?xml version="1.0"?> <svg width="508.204" height="141.732" xmlns="http://www.w3.org/2000/svg"> <path fill="#DB362D" d="M91.991,104.699c1.576,5.961,4.119,8.266,8.613,8.266c4.659,0,7.102-2.799,7.102-8.266V3.2h29.184v101.499 c0,14.307-1.856,20.506-9.11,27.762c-5.228,5.229-14.871,9.271-27.047,9.271c-9.837,0-19.25-3.256-25.253-9.27 c-5.263-5.273-8.154-10.689-12.672-27.764L44.9,37.033c-1.577-5.961-4.119-8.265-8.613-8.265c-4.66,0-7.103,2.798-7.103,8.265 v101.5H0v-101.5C0,22.727,1.857,16.527,9.111,9.271C14.337,4.044,23.981,0,36.158,0c9.837,0,19.25,3.257,25.253,9.27 c5.263,5.273,8.154,10.689,12.672,27.764L91.991,104.699z"/>  <path fill="#DB362D" d="M478.038,138.533L444.334,33.096c-0.372-1.164-0.723-2.152-1.263-2.811 c-0.926-1.127-2.207-1.719-3.931-1.719c-1.723,0-3.004,0.592-3.931,1.719c-0.539,0.658-0.891,1.646-1.262,2.811l-33.703,105.437 h-30.167l36.815-115.177c1.918-6,4.66-11.094,8.139-14.488C421.002,3.047,428.038,0,439.141,0s18.14,3.047,24.109,8.867 c3.479,3.395,6.221,8.488,8.14,14.488l36.814,115.177H478.038z"/>  <path fill="#DB362D" d="M328.878,138.533c19.12,0,28.446-4.062,35.814-11.389c8.153-8.105,12.053-16.973,12.053-30.213 c0-11.699-4.283-22.535-10.804-29.019c-8.526-8.479-19.116-11.151-36.384-11.151L305.37,56.76c-9.242,0-12.925-1.117-15.839-3.98 c-2.001-1.964-2.939-4.885-2.939-8.328c0-3.559,0.857-7.074,3.303-9.475c2.171-2.131,5.13-3.109,10.816-3.109h69.903V3.2H306.05 c-19.12,0-28.445,4.063-35.814,11.389c-8.152,8.105-12.053,16.972-12.053,30.212c0,11.701,4.283,22.536,10.804,29.019 c8.527,8.479,19.116,11.152,36.384,11.152l24.188,0.002c9.242,0,12.925,1.115,15.839,3.979c2.001,1.965,2.939,4.885,2.939,8.328 c0,3.559-0.857,7.074-3.302,9.475c-2.172,2.131-5.131,3.109-10.817,3.109h-72.094l-27.651-86.509 c-1.918-6-4.66-11.094-8.139-14.488C220.363,3.047,213.327,0,202.224,0s-18.14,3.047-24.108,8.867 c-3.48,3.395-6.221,8.488-8.139,14.488l-36.815,115.177h30.166l33.704-105.437c0.372-1.164,0.723-2.152,1.263-2.811 c0.926-1.127,2.208-1.719,3.931-1.719s3.004,0.592,3.931,1.719c0.54,0.658,0.891,1.646,1.262,2.811l33.704,105.437H328.878z"/> </svg>`
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.IsExpanded = true
	__Library__00000000_.IsRootLibrary = true

	__Note__00000000_.Name = `CFT ended in march 2025`
	__Note__00000000_.ComputedPrefix = `1`
	__Note__00000000_.IsExpanded = false

	__Note__00000001_.Name = `A thorough review of the STAR report is advised (p14)
`
	__Note__00000001_.ComputedPrefix = `2`
	__Note__00000001_.IsExpanded = false

	__Note__00000002_.Name = `NASA utilized a firm fixed price contracting type for CCtCap. 

This was a significant shift from the cost-plus contracting for traditional NASA builds of developmental vehicles. These shifts signified that CCP was not only positioned to be an innovative, first-of-its kind program for NASA, but how it interacted with new and traditional space flight industry providers was setup to be significantly distinct and different.  `
	__Note__00000002_.ComputedPrefix = `3`
	__Note__00000002_.IsExpanded = false

	__Note__00000003_.Name = `The CCP 1100 series of requirements were deliberately written at a higher-level, leaving room for provider innovation but there was also room for incorrect/inadequate interpretation by the providers. `
	__Note__00000003_.ComputedPrefix = `4`
	__Note__00000003_.IsExpanded = false

	__Note__00000004_.Name = `The Commercial Provider focused on meeting contractual requirement language resulting in insufficient demonstration across the components/system and ground/flight. `
	__Note__00000004_.ComputedPrefix = `5`
	__Note__00000004_.IsExpanded = false

	__Note__00000005_.Name = `Suppliers’ build quality/variability issues can be hard to exonerate for service modules, which is hardware that is disposed of during re-entry. `
	__Note__00000005_.ComputedPrefix = `6`
	__Note__00000005_.IsExpanded = false

	__Note__00000006_.Name = `The Aerojet Rocketdyne (AR) thermal model included the effects of jet firings, but these effects were not validated by ground testing.

Boeing thermal model did not include the effects of jet firings before CFT. `
	__Note__00000006_.ComputedPrefix = `7`
	__Note__00000006_.IsExpanded = false

	__Note__00000007_.Name = `The thruster performance from OFT1 & OFT2 experienced greater than expected temperatures and continuing to operate lead to a false sense of security of the thruster capability/performance. `
	__Note__00000007_.ComputedPrefix = `8`
	__Note__00000007_.IsExpanded = false

	__Note__00000008_.Name = `Flight instrumentation locations for thermal sensors were limited and in different locations than the locations for RCS Thruster ground firings. `
	__Note__00000008_.ComputedPrefix = `9`
	__Note__00000008_.IsExpanded = false

	__Note__00000009_.Name = `OFT1 & OFT2 investigations did not include RCS/OMAC thruster firings and fault trees were not validated through subsequent ground testing. `
	__Note__00000009_.ComputedPrefix = `10`
	__Note__00000009_.IsExpanded = false

	__Note__00000010_.Name = `For OFT2, NASA/Boeing did not have tools to measure thruster degradation, simply treated the thruster as failed/operational. `
	__Note__00000010_.ComputedPrefix = `11`
	__Note__00000010_.IsExpanded = false

	__Note__00000011_.Name = `Pc (Chamber Pressure).To know if a thruster is actually firing, the spacecraft's computer looks at the pressure sensor inside the combustion chamber (the Pc telemetry). If the pressure shoots up to the expected level, it means combustion is happening and the thruster is pushing. If the pressure stays low or at zero, it assumes the thruster failed.`
	__Note__00000011_.ComputedPrefix = `12`
	__Note__00000011_.IsExpanded = false

	__Note__00000012_.Name = `Pc transducer failed but the engine was able to restart. Boing confidence in the harxware robustness`
	__Note__00000012_.ComputedPrefix = `13`
	__Note__00000012_.IsExpanded = false

	__Note__00000013_.Name = `The leading theory for the proximate cause of the CM RCS failure during the CFT is the formulation of carbazic acid, which corrodes stainless steel. The reaction between carbazic acid and stainless steel creates corrosion particulates in the thruster propellant valve, preventing it from opening. `
	__Note__00000013_.ComputedPrefix = `14`
	__Note__00000013_.IsExpanded = false

	__Note__00000014_.Name = `Thruster health is monitored via chamber pressure and Fuel/Oxydizer injector temperatures`
	__Note__00000014_.ComputedPrefix = `15`
	__Note__00000014_.IsExpanded = false

	__Note__00000015_.Name = `Commanded pulse lengths are neither recorded nor downlinked near real time with telemetry

This complicates efforts to understand what went wrong with SM RCS jets when data suggests that pulse length is correlated to observed soakback temperatures.  `
	__Note__00000015_.ComputedPrefix = `16`
	__Note__00000015_.IsExpanded = false

	__Note__00000016_.Name = `Many thruster pulses are less  FDIR in the  The FMCs see chamber pressure over a data bus from the  but the chamber pressure is only sent to the ground with a sample rate for recording every  limiting insight into thruster performance. This does not meet Nyquist Criterion for capturing the chamber pressure signal. As a result, aliasing effects may occur for pulses shorter than  where high-frequency components of the pressure signal are misrepresented or lost, further complicating accurate reconstruction of thruster behavior.  `
	__Note__00000016_.ComputedPrefix = `17`
	__Note__00000016_.IsExpanded = false

	__Note__00000017_.Name = ` S2A2 failed at 9050 m (GMT 14:00) and B1A3 failed at 526m (GMT 14:57).`
	__Note__00000017_.ComputedPrefix = `18`
	__Note__00000017_.IsExpanded = false

	__Note__00000018_.Name = `At the conclusion of the hotfire, when thruster is demonstrated, a jet is reenabled with jet-fail FDIR inhibited. Through this sequence, four of five jets were recovered. Spacecraft 6DOF control was reestablished in all axes. T

he spacecraft was moded back to auto and the rendezvous/dock was completed.  While the flight rules allow an SM RCS thruster to be used with jet-fail off FDIR inhibited, doing so removes a hazard control.

Whaou.`
	__Note__00000018_.ComputedPrefix = `19`
	__Note__00000018_.IsExpanded = false

	__Note__00000019_.Name = ` degraded performance was still observed. This testing demonstrated the issue was not purely transient.

It should have been done before ?!

 In addition to demonstrating transient degradation during high heating, the test program demonstrated an unrecoverable cumulative degradation that was potentially indicative of the observed performance of CFT. `
	__Note__00000019_.ComputedPrefix = `20`
	__Note__00000019_.IsExpanded = false

	__Note__00000020_.Name = `B1.3.3.1.1.1 Heat identifies that excessive heating could cause the poppet to extrude, reducing flow. Poppet extrusion was observed in the post-test tear down of the WSTF (White Sands Test Facility) thruster unit.

The initial heating was caused by the internal doghouse temperatures (addressed in C4.1 Thermal) and the follow-on heating induced by B1.3.3.1.4.2 RCS Thermal Soakback. This caused excessive temperatures beyond the capability of the thruster softgoods, causing the Teflon poppet to extrude. The evidence supports this being a contributor to the thruster fail-offs.`
	__Note__00000020_.ComputedPrefix = `21`
	__Note__00000020_.IsExpanded = false

	__Note__00000021_.Name = `B1.3.3.1.4.2 RCS Thermal Soakback identifies the possibility that structure of the RCS thruster
retains heat, and the heat is unable to be dissipated from the firings of the thrusters, causing the
propellants to heat rapidly. `
	__Note__00000021_.ComputedPrefix = `22`
	__Note__00000021_.IsExpanded = false

	__Note__00000022_.Name = `Intermediate Cause 5: Inadequate Thruster Thermal Models

In May 2019, the PCB accepted the “SM doghouse thermal model (PROP-06) not in critical model
Database” risk, which was generated in 2017. As identified by Passive Thermal Control System
(PTCS), “there was a gap in Boeing thermal modeling. Specifically, the internal componentry within
each of the doghouses isn’t modeled in THERM-16. This area is modeled by Aerojet-Rocketdyne (AR)
per contract with Boeing PROP. The contract only furnishes the results and not the model that
produces the results. a.) AR model needs to undergoes accreditation. There is no NASA insight into
the validity of the results and therefore verification effort is compromised.”

The Therm-11a model did not incorporate OMAC firings, or initial heating that could occur from the
doghouse environment, which is critical given that the aft thrusters have the longest feed line in
comparison to every other thruster.

There was no thermal model available to adequately understand the full impacts of the thermal
environment upon the RCS thrusters, specifically the aft thrusters with the longest feed line tube,
and very close proximity to the OMAC thrusters`
	__Note__00000022_.ComputedPrefix = `23`
	__Note__00000022_.IsExpanded = false

	__Note__00000023_.Name = `Intermediate Cause 6: Insufficient Thruster Qualification

SM RCS Thruster Qualification did not cover the flight envelope for temperature and duty cycle (TLYF).`
	__Note__00000023_.ComputedPrefix = `24`
	__Note__00000023_.IsExpanded = false

	__Note__00000024_.Name = `Multiple groups, prior to CTF, discussed the lack of mission representative operational duty cycle testing for SM RCS engines.. This concern is one that has been tracked since prior to OFT1 and is specifically reviewed in this intermediate cause because it is widely accepted that operational duty cycle is a primary driver of hardware temperature. NASA Engineering and the CCP Spacecraft Office have consistently identified this qualification gap and recommended additional testing, however it was not incorporated into the pursued plans for risk assessments between flights.

OFT1: The SM RCS Qual Gaps risk was accepted in September 2019 (PCB-19-383), Boeing
RCS/OMAC (Engine) Qual Issues for OFT.

CFT: The qualification gap risk was addressed in April 2023

. The program directive issued at that time resolved to accept
the qualification gap for duty cycles for CFT only. This acceptance was encompassed within the
previously approved elevated 2x5 risk level from PCB-23-053, which was documented during the
IFA closure for OFT-2 related to valve/injector assembly temperature measurements exceeding
qualification limits. These temperature exceedances were categorized as an Unexplained Anomaly
(UA).
Starliner Qualification Risk History
OFT-1 (Orbital Flight Test 1)
The Service Module (SM) RCS Qual Gaps risk was formally accepted in September 2019 (PCB-19-383) regarding Boeing RCS/OMAC (Engine) qualification issues for the flight.

OFT-2 (Orbital Flight Test 2)
During PCB-20-404, addressing the SM RCS Jet Failure, the program decided to use the SM RCS "as is." It was noted that the OFT Mission Elapsed Time (MET) anomaly actually demonstrated the robustness of the RCS thruster design despite long durations of excessive usage. While design modifications were evaluated for the leading theory behind the B2R3 failure, the CCP SDRT report highlighted this as a missed opportunity to further investigate SM RCS thruster failure modes. This is discussed further in later sections regarding organizational factors and anomaly resolution.

CFT (Crew Flight Test)
The qualification gap risk was addressed in April 2023 during PCB-23-100 (CFT: Starliner Prop, OMAC/RCS Hot-Fire Qual Gaps). The program directive resolved to accept the qualification gap for duty cycles for the CFT only. This acceptance was encompassed within a previously approved elevated 2x5 risk level (PCB-23-053), documented during the In-Flight Anomaly (IFA) closure for OFT-2.

The OFT-2 IFA related to valve/injector assembly temperature measurements exceeding qualification limits, which were categorized as an Unexplained Anomaly (UA). The directive acknowledged that these exceedances were attributed to specific operational duty cycles. Consequently, the thermal risk was accepted as a 2x5 risk for potential jet failure. Despite this, engineering teams and the Engineering Review Board (ERB-23-0045-R2) emphasized the importance of testing SM RCS thrusters using flight-like duty cycles.`
	__Note__00000024_.ComputedPrefix = `25`
	__Note__00000024_.IsExpanded = false

	__Note__00000025_.Name = `Organizational Factor 1: Insufficient Anomaly Resolution Process Multiple RCS thruster failures occurred on OFT 1 and OFT 2. While fault trees were developed, due diligence to reach direct and root cause was not performed and adequate block closure rationale was often incomplete.

Note : Anomaly Reports (AR) were closed.

Organizational Factor 2: Mischaracterization of Risk in thruster Qual Gaps leads to Flawed Flight Rationale. The thermal risk for Starliner’s SM RCS thrusters was mischaracterized due to reliance on simple models and inadequate qualification testing that failed to replicate mission-representative conditions. Despite clear evidence of thermal soakback and temperature exceedances, risk acceptance proceeded without resolving key environment and duty cycle concerns.`
	__Note__00000025_.ComputedPrefix = `26`
	__Note__00000025_.IsExpanded = false

	__NoteProductShape__00000000_.Name = `A thorough review of the STAR report is advised to  Starliner Tests and Anomalies Review (STAR) Investigation Report`
	__NoteProductShape__00000000_.StartRatio = 0.500000
	__NoteProductShape__00000000_.EndRatio = 0.500000
	__NoteProductShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000000_.CornerOffsetRatio = 1.680000
	__NoteProductShape__00000000_.IsHidden = false

	__NoteProductShape__00000001_.Name = `NASA utilized a firm fixed price contracting type for CCtCap. This was a significant shift from the cost-plus contracting for traditional NASA builds of developmental vehicles. These shifts signified that CCP was not only positioned to be an innovative, first-of-its kind program for NASA, but how it interacted with new and traditional space flight industry providers was setup to be significantly distinct and different.   to Commercial Crew Transportation Capability (CCtCap).`
	__NoteProductShape__00000001_.StartRatio = 0.500000
	__NoteProductShape__00000001_.EndRatio = 0.500000
	__NoteProductShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000001_.CornerOffsetRatio = 1.680000
	__NoteProductShape__00000001_.IsHidden = false

	__NoteProductShape__00000002_.Name = ` to CCP Requirements`
	__NoteProductShape__00000002_.StartRatio = 0.500000
	__NoteProductShape__00000002_.EndRatio = 0.500000
	__NoteProductShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000002_.CornerOffsetRatio = 1.680000
	__NoteProductShape__00000002_.IsHidden = false

	__NoteProductShape__00000003_.Name = `The Commercial Provider focused on meeting contractual requirement language resulting in insufficient demonstration across the components/system and ground/flight.  to CCP Requirements`
	__NoteProductShape__00000003_.StartRatio = 0.500000
	__NoteProductShape__00000003_.EndRatio = 0.500000
	__NoteProductShape__00000003_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000003_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000003_.CornerOffsetRatio = 1.680000
	__NoteProductShape__00000003_.IsHidden = false

	__NoteProductShape__00000004_.Name = `Suppliers’ build quality/variability issues can be hard to exonerate for service modules, which is hardware that is disposed of during re-entry.  to CM RCS Thrusters`
	__NoteProductShape__00000004_.StartRatio = 0.500000
	__NoteProductShape__00000004_.EndRatio = 0.500000
	__NoteProductShape__00000004_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000004_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000004_.CornerOffsetRatio = 1.680000
	__NoteProductShape__00000004_.IsHidden = false

	__NoteProductShape__00000005_.Name = `Suppliers’ build quality/variability issues can be hard to exonerate for service modules, which is hardware that is disposed of during re-entry.  to SM RCS Thrusters`
	__NoteProductShape__00000005_.StartRatio = 0.500000
	__NoteProductShape__00000005_.EndRatio = 0.500000
	__NoteProductShape__00000005_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000005_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000005_.CornerOffsetRatio = 1.680000
	__NoteProductShape__00000005_.IsHidden = false

	__NoteProductShape__00000006_.Name = `The Aerojet Rocketdyne (AR) thermal model included the effects of jet firings, but these effects were not validated by ground testing.

Boeing thermal model did not include the effects of jet firings before CFT.  to 28 SM RCS Thrusters`
	__NoteProductShape__00000006_.StartRatio = 0.500000
	__NoteProductShape__00000006_.EndRatio = 0.500000
	__NoteProductShape__00000006_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000006_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000006_.CornerOffsetRatio = 1.680000
	__NoteProductShape__00000006_.IsHidden = false

	__NoteProductShape__00000007_.Name = `The thruster performance from OFT1 & OFT2 experienced greater than expected temperatures and continuing to operate lead to a false sense of security of the thruster capability/performance.  to 28 SM RCS Thrusters`
	__NoteProductShape__00000007_.StartRatio = 0.500000
	__NoteProductShape__00000007_.EndRatio = 0.500000
	__NoteProductShape__00000007_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000007_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000007_.CornerOffsetRatio = 1.680000
	__NoteProductShape__00000007_.IsHidden = false

	__NoteProductShape__00000008_.Name = `Flight instrumentation locations for thermal sensors were limited and in different locations than the locations for RCS Thruster ground firings.  to Thermal Sensors`
	__NoteProductShape__00000008_.StartRatio = 0.500000
	__NoteProductShape__00000008_.EndRatio = 0.500000
	__NoteProductShape__00000008_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000008_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000008_.CornerOffsetRatio = 1.680000
	__NoteProductShape__00000008_.IsHidden = false

	__NoteProductShape__00000009_.Name = `Flight instrumentation locations for thermal sensors were limited and in different locations than the locations for RCS Thruster ground firings.  to 28 SM RCS Thrusters`
	__NoteProductShape__00000009_.StartRatio = 0.500000
	__NoteProductShape__00000009_.EndRatio = 0.500000
	__NoteProductShape__00000009_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000009_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000009_.CornerOffsetRatio = 1.680000
	__NoteProductShape__00000009_.IsHidden = false

	__NoteProductShape__00000010_.Name = `Flight instrumentation locations for thermal sensors were limited and in different locations than the locations for RCS Thruster ground firings.  to 12 CM RCS Thrusters`
	__NoteProductShape__00000010_.StartRatio = 0.500000
	__NoteProductShape__00000010_.EndRatio = 0.500000
	__NoteProductShape__00000010_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000010_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000010_.CornerOffsetRatio = 1.680000
	__NoteProductShape__00000010_.IsHidden = false

	__NoteProductShape__00000011_.Name = `OFT1 & OFT2 investigations did not include RCS/OMAC thruster firings and fault trees were not validated through subsequent ground testing.  to 28 SM RCS Thrusters (2800 lbf)`
	__NoteProductShape__00000011_.StartRatio = 0.500000
	__NoteProductShape__00000011_.EndRatio = 0.500000
	__NoteProductShape__00000011_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000011_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000011_.CornerOffsetRatio = 1.680000
	__NoteProductShape__00000011_.IsHidden = false

	__NoteProductShape__00000012_.Name = `For OFT2, NASA/Boeing did not have tools to measure thruster degradation, simply treated the thruster as failed/operational.  to 28 SM RCS Thrusters (2800 lbf)`
	__NoteProductShape__00000012_.StartRatio = 0.500000
	__NoteProductShape__00000012_.EndRatio = 0.500000
	__NoteProductShape__00000012_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000012_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000012_.CornerOffsetRatio = 1.680000
	__NoteProductShape__00000012_.IsHidden = false

	__NoteProductShape__00000013_.Name = `Pc (Chamber Pressure) to 12 CM RCS Thrusters (1200 lbf)`
	__NoteProductShape__00000013_.StartRatio = 0.500000
	__NoteProductShape__00000013_.EndRatio = 0.500000
	__NoteProductShape__00000013_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000013_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000013_.CornerOffsetRatio = 1.680000
	__NoteProductShape__00000013_.IsHidden = false

	__NoteProductShape__00000014_.Name = `Pc (Chamber Pressure).To know if a thruster is actually firing, the spacecraft's computer looks at the pressure sensor inside the combustion chamber (the Pc telemetry). If the pressure shoots up to the expected level, it means combustion is happening and the thruster is pushing. If the pressure stays low or at zero, it assumes the thruster failed. to Transducer / pressure sensor`
	__NoteProductShape__00000014_.StartRatio = 0.500000
	__NoteProductShape__00000014_.EndRatio = 0.500000
	__NoteProductShape__00000014_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000014_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000014_.CornerOffsetRatio = 1.680000
	__NoteProductShape__00000014_.IsHidden = false

	__NoteProductShape__00000015_.Name = `Pc transducer failed but the engine was able to restart. Boing confidence in the harxware robustness to Transducer / pressure sensor (Pc)`
	__NoteProductShape__00000015_.StartRatio = 0.500000
	__NoteProductShape__00000015_.EndRatio = 0.500000
	__NoteProductShape__00000015_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000015_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000015_.CornerOffsetRatio = 1.680000
	__NoteProductShape__00000015_.IsHidden = false

	__NoteProductShape__00000016_.Name = `The leading theory for the proximate cause of the CM RCS failure during the CFT is the formulation of carbazic acid, which corrodes stainless steel. The reaction between carbazic acid and stainless steel creates corrosion particulates in the thruster propellant valve, preventing it from opening.  to 4.4   Analysis: CM RCS Jet Failure`
	__NoteProductShape__00000016_.StartRatio = 0.500000
	__NoteProductShape__00000016_.EndRatio = 0.500000
	__NoteProductShape__00000016_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000016_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000016_.CornerOffsetRatio = 1.680000
	__NoteProductShape__00000016_.IsHidden = false

	__NoteProductShape__00000017_.Name = `Thruster health is monitored via chamber pressure and Fuel/Oxydizer injector temperatures to 4.4   Analysis: CM RCS Jet Failure`
	__NoteProductShape__00000017_.StartRatio = 0.500000
	__NoteProductShape__00000017_.EndRatio = 0.500000
	__NoteProductShape__00000017_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000017_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000017_.CornerOffsetRatio = 1.680000
	__NoteProductShape__00000017_.IsHidden = false

	__NoteProductShape__00000018_.Name = `Commanded pulse lengths are neither recorded nor downlinked near real time with telemetry to 4.5.1 Description of the system`
	__NoteProductShape__00000018_.StartRatio = 0.500000
	__NoteProductShape__00000018_.EndRatio = 0.500000
	__NoteProductShape__00000018_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000018_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000018_.CornerOffsetRatio = 1.680000
	__NoteProductShape__00000018_.IsHidden = false

	__NoteProductShape__00000019_.Name = `At the conclusion of the hotfire, when thruster is demonstrated, a jet is reenabled with jet-fail FDIR inhibited. Through this sequence, four of five jets were recovered. Spacecraft 6DOF control was reestablished in all axes. The spacecraft was moded back to auto and the rendezvous/dock was completed.  While the flight rules allow an SM RCS thruster to be used with jet-fail off FDIR inhibited, doing so removes a hazard control to 4.5.2   Description of Events and Timeline `
	__NoteProductShape__00000019_.StartRatio = 0.500000
	__NoteProductShape__00000019_.EndRatio = 0.500000
	__NoteProductShape__00000019_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000019_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000019_.CornerOffsetRatio = 1.680000
	__NoteProductShape__00000019_.IsHidden = false

	__NoteProductShape__00000020_.Name = `Many thruster pulses are less  FDIR in the  The FMCs see chamber pressure over a data bus from the  but the chamber pressure is only sent to the ground with a sample rate for recording every  limiting insight into thruster performance. This does not meet Nyquist Criterion for capturing the chamber pressure signal. As a result, aliasing effects may occur for pulses shorter than  where high-frequency components of the pressure signal are misrepresented or lost, further complicating accurate reconstruction of thruster behavior.   to 4.5.2   Description of Events and Timeline `
	__NoteProductShape__00000020_.StartRatio = 0.500000
	__NoteProductShape__00000020_.EndRatio = 0.500000
	__NoteProductShape__00000020_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000020_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000020_.CornerOffsetRatio = 1.680000
	__NoteProductShape__00000020_.IsHidden = false

	__NoteProductShape__00000021_.Name = ` degraded performance was still observed. This testing demonstrated the issue was not purely transient. to 4.5.4   Starliner Engine Testing at White Sands Test Facility (WSTF) during CFT`
	__NoteProductShape__00000021_.StartRatio = 0.500000
	__NoteProductShape__00000021_.EndRatio = 0.500000
	__NoteProductShape__00000021_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000021_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000021_.CornerOffsetRatio = 1.680000
	__NoteProductShape__00000021_.IsHidden = false

	__NoteProductShape__00000022_.Name = `B1.3.3.1.1.1 Heat identifies that excessive heating could cause the poppet to extrude, reducing flow. Poppet extrusion was observed in the post-test tear down of the WSTF (White Sands Test Facility) thruster unit.

The initial heating was caused by the internal doghouse temperatures (addressed in C4.1 Thermal) and the follow-on heating induced by B1.3.3.1.4.2 RCS Thermal Soakback. This caused excessive temperatures beyond the capability of the thruster softgoods, causing the Teflon poppet to extrude. The evidence supports this being a contributor to the thruster fail-offs. to 4.5.5 Fault Tree`
	__NoteProductShape__00000022_.StartRatio = 0.500000
	__NoteProductShape__00000022_.EndRatio = 0.500000
	__NoteProductShape__00000022_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000022_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000022_.CornerOffsetRatio = 1.680000
	__NoteProductShape__00000022_.IsHidden = false

	__NoteProductShape__00000023_.Name = `B1.3.3.1.4.2 RCS Thermal Soakback identifies the possibility that structure of the RCS thruster
retains heat, and the heat is unable to be dissipated from the firings of the thrusters, causing the
propellants to heat rapidly.  to 4.5.5 Fault Tree`
	__NoteProductShape__00000023_.StartRatio = 0.500000
	__NoteProductShape__00000023_.EndRatio = 0.500000
	__NoteProductShape__00000023_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000023_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000023_.CornerOffsetRatio = 1.680000
	__NoteProductShape__00000023_.IsHidden = false

	__NoteProductShape__00000024_.Name = `In May 2019, the PCB accepted the “SM doghouse thermal model (PROP-06) not in critical model
Database” risk, which was generated in 2017. As identified by Passive Thermal Control System
(PTCS), “there was a gap in Boeing thermal modeling. Specifically, the internal componentry within
each of the doghouses isn’t modeled in THERM-16. This area is modeled by Aerojet-Rocketdyne (AR)
per contract with Boeing PROP. The contract only furnishes the results and not the model that
produces the results. a.) AR model needs to undergoes accreditation. There is no NASA insight into
the validity of the results and therefore verification effort is compromised.” to Intermediate Cause 5: Inadequate Thruster Thermal Models Inadequate thermal modelling caused insufficient scrutiny for the thermal environment, leading to excessive heating from RCS thermal soakback and integrated heating from OMACs.`
	__NoteProductShape__00000024_.StartRatio = 0.500000
	__NoteProductShape__00000024_.EndRatio = 0.500000
	__NoteProductShape__00000024_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000024_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000024_.CornerOffsetRatio = 1.680000
	__NoteProductShape__00000024_.IsHidden = false

	__NoteProductShape__00000025_.Name = `Intermediate Cause 6: Insufficient Thruster Qualification

SM RCS Thruster Qualification did not cover the flight envelope for temperature and duty cycle (TLYF). to Intermediate Cause 5: Inadequate Thruster Thermal Models Inadequate thermal modelling caused insufficient scrutiny for the thermal environment, leading to excessive heating from RCS thermal soakback and integrated heating from OMACs.`
	__NoteProductShape__00000025_.StartRatio = 0.500000
	__NoteProductShape__00000025_.EndRatio = 0.500000
	__NoteProductShape__00000025_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000025_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000025_.CornerOffsetRatio = 1.680000
	__NoteProductShape__00000025_.IsHidden = false

	__NoteProductShape__00000026_.Name = `Multiple groups, prior to CTF, discussed the lack of mission representative operational duty cycle testing for SM RCS engines.. This concern is one that has been tracked since prior to OFT1 and is specifically reviewed in this intermediate cause because it is widely accepted that operational duty cycle is a primary driver of hardware temperature. NASA Engineering and the CCP Spacecraft Office have consistently identified this qualification gap and recommended additional testing, however it was not incorporated into the pursued plans for risk assessments between flights. to Intermediate Cause 5: Inadequate Thruster Thermal Models Inadequate thermal modelling caused insufficient scrutiny for the thermal environment, leading to excessive heating from RCS thermal soakback and integrated heating from OMACs.`
	__NoteProductShape__00000026_.StartRatio = 0.500000
	__NoteProductShape__00000026_.EndRatio = 0.500000
	__NoteProductShape__00000026_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000026_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000026_.CornerOffsetRatio = 1.680000
	__NoteProductShape__00000026_.IsHidden = false

	__NoteShape__00000001_.Name = `-PIT focus`
	__NoteShape__00000001_.X = 778.606649
	__NoteShape__00000001_.Y = 81.204707
	__NoteShape__00000001_.Width = 306.000000
	__NoteShape__00000001_.Height = 70.000000
	__NoteShape__00000001_.IsHidden = false

	__NoteShape__00000002_.Name = `-PBS`
	__NoteShape__00000002_.X = 566.290428
	__NoteShape__00000002_.Y = 403.921876
	__NoteShape__00000002_.Width = 294.000000
	__NoteShape__00000002_.Height = 266.999985
	__NoteShape__00000002_.IsHidden = false

	__NoteShape__00000003_.Name = `-PIT Report`
	__NoteShape__00000003_.X = 1285.673362
	__NoteShape__00000003_.Y = 221.503075
	__NoteShape__00000003_.Width = 250.000000
	__NoteShape__00000003_.Height = 177.000000
	__NoteShape__00000003_.IsHidden = false

	__NoteShape__00000004_.Name = `-PIT Report`
	__NoteShape__00000004_.X = 744.720613
	__NoteShape__00000004_.Y = 305.622045
	__NoteShape__00000004_.Width = 250.000000
	__NoteShape__00000004_.Height = 143.999939
	__NoteShape__00000004_.IsHidden = false

	__NoteShape__00000005_.Name = `-RCS PBS`
	__NoteShape__00000005_.X = 334.089724
	__NoteShape__00000005_.Y = 779.092030
	__NoteShape__00000005_.Width = 250.000000
	__NoteShape__00000005_.Height = 194.999939
	__NoteShape__00000005_.IsHidden = false

	__NoteShape__00000006_.Name = `-RCS PBS`
	__NoteShape__00000006_.X = 1219.153628
	__NoteShape__00000006_.Y = 777.461448
	__NoteShape__00000006_.Width = 250.000000
	__NoteShape__00000006_.Height = 196.000000
	__NoteShape__00000006_.IsHidden = false

	__NoteShape__00000007_.Name = `-RCS PBS`
	__NoteShape__00000007_.X = 596.510371
	__NoteShape__00000007_.Y = 678.471309
	__NoteShape__00000007_.Width = 269.000000
	__NoteShape__00000007_.Height = 181.999939
	__NoteShape__00000007_.IsHidden = false

	__NoteShape__00000008_.Name = `-RCS PBS`
	__NoteShape__00000008_.X = 917.375080
	__NoteShape__00000008_.Y = 517.315922
	__NoteShape__00000008_.Width = 250.000000
	__NoteShape__00000008_.Height = 144.000000
	__NoteShape__00000008_.IsHidden = false

	__NoteShape__00000009_.Name = `-RCS PBS`
	__NoteShape__00000009_.X = 880.446589
	__NoteShape__00000009_.Y = 755.302818
	__NoteShape__00000009_.Width = 250.000000
	__NoteShape__00000009_.Height = 154.000061
	__NoteShape__00000009_.IsHidden = false

	__NoteShape__00000010_.Name = `-RCS PBS`
	__NoteShape__00000010_.X = 861.189303
	__NoteShape__00000010_.Y = 1040.200971
	__NoteShape__00000010_.Width = 250.000000
	__NoteShape__00000010_.Height = 114.000000
	__NoteShape__00000010_.IsHidden = false

	__NoteShape__00000011_.Name = `-RCS PBS`
	__NoteShape__00000011_.X = 45.256122
	__NoteShape__00000011_.Y = 950.529205
	__NoteShape__00000011_.Width = 250.000000
	__NoteShape__00000011_.Height = 227.000061
	__NoteShape__00000011_.IsHidden = false

	__NoteShape__00000012_.Name = `-RCS PBS`
	__NoteShape__00000012_.X = 565.566920
	__NoteShape__00000012_.Y = 924.052255
	__NoteShape__00000012_.Width = 250.000000
	__NoteShape__00000012_.Height = 121.000000
	__NoteShape__00000012_.IsHidden = false

	__NoteShape__00000013_.Name = `-PIT Report`
	__NoteShape__00000013_.X = 1150.711246
	__NoteShape__00000013_.Y = 600.447807
	__NoteShape__00000013_.Width = 250.000000
	__NoteShape__00000013_.Height = 271.000000
	__NoteShape__00000013_.IsHidden = false

	__NoteShape__00000014_.Name = `-PIT Report`
	__NoteShape__00000014_.X = 1335.901110
	__NoteShape__00000014_.Y = 996.520214
	__NoteShape__00000014_.Width = 250.000000
	__NoteShape__00000014_.Height = 142.000000
	__NoteShape__00000014_.IsHidden = false

	__NoteShape__00000015_.Name = `-PIT Report`
	__NoteShape__00000015_.X = 95.642801
	__NoteShape__00000015_.Y = 1355.855478
	__NoteShape__00000015_.Width = 250.000000
	__NoteShape__00000015_.Height = 193.000000
	__NoteShape__00000015_.IsHidden = false

	__NoteShape__00000016_.Name = `-PIT Report`
	__NoteShape__00000016_.X = 425.148282
	__NoteShape__00000016_.Y = 1430.634730
	__NoteShape__00000016_.Width = 314.000000
	__NoteShape__00000016_.Height = 278.000000
	__NoteShape__00000016_.IsHidden = false

	__NoteShape__00000017_.Name = `-WBS`
	__NoteShape__00000017_.X = 591.638035
	__NoteShape__00000017_.Y = 652.450426
	__NoteShape__00000017_.Width = 292.000000
	__NoteShape__00000017_.Height = 70.000000
	__NoteShape__00000017_.IsHidden = false

	__NoteShape__00000018_.Name = `-PIT Report`
	__NoteShape__00000018_.X = 753.803011
	__NoteShape__00000018_.Y = 1447.804082
	__NoteShape__00000018_.Width = 250.000000
	__NoteShape__00000018_.Height = 333.000000
	__NoteShape__00000018_.IsHidden = false

	__NoteShape__00000019_.Name = `-PIT Report`
	__NoteShape__00000019_.X = 1024.260622
	__NoteShape__00000019_.Y = 1594.529677
	__NoteShape__00000019_.Width = 250.000000
	__NoteShape__00000019_.Height = 280.000000
	__NoteShape__00000019_.IsHidden = false

	__NoteShape__00000020_.Name = `-PIT Report`
	__NoteShape__00000020_.X = 1325.045052
	__NoteShape__00000020_.Y = 1477.461652
	__NoteShape__00000020_.Width = 250.000000
	__NoteShape__00000020_.Height = 383.000000
	__NoteShape__00000020_.IsHidden = false

	__NoteShape__00000021_.Name = `-PIT Report`
	__NoteShape__00000021_.X = 1009.364672
	__NoteShape__00000021_.Y = 1927.551182
	__NoteShape__00000021_.Width = 250.000000
	__NoteShape__00000021_.Height = 188.000000
	__NoteShape__00000021_.IsHidden = false

	__NoteShape__00000022_.Name = `-PIT Report`
	__NoteShape__00000022_.X = 409.341567
	__NoteShape__00000022_.Y = 1822.743943
	__NoteShape__00000022_.Width = 551.000000
	__NoteShape__00000022_.Height = 496.000000
	__NoteShape__00000022_.IsHidden = false

	__NoteShape__00000023_.Name = `-PIT Report`
	__NoteShape__00000023_.X = 258.414750
	__NoteShape__00000023_.Y = 2019.668077
	__NoteShape__00000023_.Width = 250.000000
	__NoteShape__00000023_.Height = 172.000000
	__NoteShape__00000023_.IsHidden = false

	__NoteShape__00000024_.Name = `-PIT Report`
	__NoteShape__00000024_.X = 177.291855
	__NoteShape__00000024_.Y = 2342.855372
	__NoteShape__00000024_.Width = 1451.000000
	__NoteShape__00000024_.Height = 594.000000
	__NoteShape__00000024_.IsHidden = false

	__NoteShape__00000025_.Name = `-PIT Report`
	__NoteShape__00000025_.X = 916.887483
	__NoteShape__00000025_.Y = 2959.225340
	__NoteShape__00000025_.Width = 723.000000
	__NoteShape__00000025_.Height = 289.000000
	__NoteShape__00000025_.IsHidden = false

	__NoteTaskShape__00000001_.Name = ` S2A2 failed at 9050 m (GMT 14:00) and B1A3 failed at 526m (GMT 14:57). to ISS Approach`
	__NoteTaskShape__00000001_.StartRatio = 0.500000
	__NoteTaskShape__00000001_.EndRatio = 0.500000
	__NoteTaskShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteTaskShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteTaskShape__00000001_.CornerOffsetRatio = 1.680000
	__NoteTaskShape__00000001_.IsHidden = false

	__Product__00000001_.Name = `Dragon`
	__Product__00000001_.ComputedPrefix = `1.1.2`
	__Product__00000001_.IsExpanded = false
	__Product__00000001_.IsImport = false
	__Product__00000001_.Description = ``
	__Product__00000001_.IsProducersNodeExpanded = false
	__Product__00000001_.IsConsumersNodeExpanded = false
	__Product__00000001_.LayoutDirection = models.Vertical

	__Product__00000002_.Name = `CST-100 Starliner`
	__Product__00000002_.ComputedPrefix = `1.1.1`
	__Product__00000002_.IsExpanded = false
	__Product__00000002_.IsImport = false
	__Product__00000002_.Description = ``
	__Product__00000002_.IsProducersNodeExpanded = false
	__Product__00000002_.IsConsumersNodeExpanded = false
	__Product__00000002_.LayoutDirection = models.Vertical

	__Product__00000003_.Name = `Reports`
	__Product__00000003_.ComputedPrefix = `2`
	__Product__00000003_.IsExpanded = false
	__Product__00000003_.IsImport = false
	__Product__00000003_.Description = ``
	__Product__00000003_.IsProducersNodeExpanded = false
	__Product__00000003_.IsConsumersNodeExpanded = false
	__Product__00000003_.LayoutDirection = models.Horizontal

	__Product__00000004_.Name = ` Starliner Tests and Anomalies Review (STAR) Investigation Report`
	__Product__00000004_.ComputedPrefix = `2.1`
	__Product__00000004_.IsExpanded = false
	__Product__00000004_.IsImport = false
	__Product__00000004_.Description = ``
	__Product__00000004_.IsProducersNodeExpanded = false
	__Product__00000004_.IsConsumersNodeExpanded = false
	__Product__00000004_.LayoutDirection = models.Vertical

	__Product__00000005_.Name = `Program Investigation Team (PIT) Reports`
	__Product__00000005_.ComputedPrefix = `2.2`
	__Product__00000005_.IsExpanded = false
	__Product__00000005_.IsImport = false
	__Product__00000005_.Description = ``
	__Product__00000005_.IsProducersNodeExpanded = false
	__Product__00000005_.IsConsumersNodeExpanded = false
	__Product__00000005_.LayoutDirection = models.Vertical

	__Product__00000006_.Name = `Commercial Crew Transportation Capability (CCtCap).`
	__Product__00000006_.ComputedPrefix = `1.1`
	__Product__00000006_.IsExpanded = false
	__Product__00000006_.IsImport = false
	__Product__00000006_.Description = ``
	__Product__00000006_.IsProducersNodeExpanded = false
	__Product__00000006_.IsConsumersNodeExpanded = false
	__Product__00000006_.LayoutDirection = models.Vertical

	__Product__00000007_.Name = `NASA Assets/Capabities`
	__Product__00000007_.ComputedPrefix = `1`
	__Product__00000007_.IsExpanded = false
	__Product__00000007_.IsImport = false
	__Product__00000007_.Description = ``
	__Product__00000007_.IsProducersNodeExpanded = false
	__Product__00000007_.IsConsumersNodeExpanded = false
	__Product__00000007_.LayoutDirection = models.Vertical

	__Product__00000008_.Name = `ISS`
	__Product__00000008_.ComputedPrefix = `1.2`
	__Product__00000008_.IsExpanded = false
	__Product__00000008_.IsImport = false
	__Product__00000008_.Description = ``
	__Product__00000008_.IsProducersNodeExpanded = false
	__Product__00000008_.IsConsumersNodeExpanded = false
	__Product__00000008_.LayoutDirection = models.Vertical

	__Product__00000009_.Name = `3 Commercial Crew Program (CCP) Background`
	__Product__00000009_.ComputedPrefix = `2.2.2`
	__Product__00000009_.IsExpanded = false
	__Product__00000009_.IsImport = false
	__Product__00000009_.Description = ``
	__Product__00000009_.IsProducersNodeExpanded = false
	__Product__00000009_.IsConsumersNodeExpanded = false
	__Product__00000009_.LayoutDirection = models.Vertical

	__Product__00000010_.Name = `CCP Requirements`
	__Product__00000010_.ComputedPrefix = `1.1.3`
	__Product__00000010_.IsExpanded = false
	__Product__00000010_.IsImport = false
	__Product__00000010_.Description = ``
	__Product__00000010_.IsProducersNodeExpanded = false
	__Product__00000010_.IsConsumersNodeExpanded = false
	__Product__00000010_.LayoutDirection = models.Vertical

	__Product__00000011_.Name = `Crew Module (CM)`
	__Product__00000011_.ComputedPrefix = `1.1.1.1`
	__Product__00000011_.IsExpanded = false
	__Product__00000011_.IsImport = false
	__Product__00000011_.Description = ``
	__Product__00000011_.IsProducersNodeExpanded = false
	__Product__00000011_.IsConsumersNodeExpanded = false
	__Product__00000011_.LayoutDirection = models.Vertical

	__Product__00000012_.Name = `CM Reaction Control System (RCS)`
	__Product__00000012_.ComputedPrefix = `1.1.1.1.1`
	__Product__00000012_.IsExpanded = false
	__Product__00000012_.IsImport = false
	__Product__00000012_.Description = ``
	__Product__00000012_.IsProducersNodeExpanded = false
	__Product__00000012_.IsConsumersNodeExpanded = false
	__Product__00000012_.LayoutDirection = models.Vertical

	__Product__00000013_.Name = `12 CM RCS Thrusters (1200 lbf)`
	__Product__00000013_.ComputedPrefix = `1.1.1.1.1.1`
	__Product__00000013_.IsExpanded = false
	__Product__00000013_.IsImport = false
	__Product__00000013_.Description = ``
	__Product__00000013_.IsProducersNodeExpanded = false
	__Product__00000013_.IsConsumersNodeExpanded = false
	__Product__00000013_.LayoutDirection = models.Vertical

	__Product__00000014_.Name = `Service Module (SM)`
	__Product__00000014_.ComputedPrefix = `1.1.1.2`
	__Product__00000014_.IsExpanded = false
	__Product__00000014_.IsImport = false
	__Product__00000014_.Description = ``
	__Product__00000014_.IsProducersNodeExpanded = false
	__Product__00000014_.IsConsumersNodeExpanded = false
	__Product__00000014_.LayoutDirection = models.Vertical

	__Product__00000015_.Name = `SM Reaction Control System (SM RCS)`
	__Product__00000015_.ComputedPrefix = `1.1.1.2.1`
	__Product__00000015_.IsExpanded = false
	__Product__00000015_.IsImport = false
	__Product__00000015_.Description = ``
	__Product__00000015_.IsProducersNodeExpanded = false
	__Product__00000015_.IsConsumersNodeExpanded = false
	__Product__00000015_.LayoutDirection = models.Vertical

	__Product__00000016_.Name = `28 SM RCS Thrusters (2800 lbf)`
	__Product__00000016_.ComputedPrefix = `1.1.1.2.1.1`
	__Product__00000016_.IsExpanded = false
	__Product__00000016_.IsImport = false
	__Product__00000016_.Description = ``
	__Product__00000016_.IsProducersNodeExpanded = false
	__Product__00000016_.IsConsumersNodeExpanded = false
	__Product__00000016_.LayoutDirection = models.Vertical

	__Product__00000017_.Name = `2 Additional Investigations `
	__Product__00000017_.ComputedPrefix = `2.2.1`
	__Product__00000017_.IsExpanded = false
	__Product__00000017_.IsImport = false
	__Product__00000017_.Description = ``
	__Product__00000017_.IsProducersNodeExpanded = false
	__Product__00000017_.IsConsumersNodeExpanded = false
	__Product__00000017_.LayoutDirection = models.Vertical

	__Product__00000018_.Name = `STAR Report Summary and Findings `
	__Product__00000018_.ComputedPrefix = `2.2.1.1`
	__Product__00000018_.IsExpanded = false
	__Product__00000018_.IsImport = false
	__Product__00000018_.Description = ``
	__Product__00000018_.IsProducersNodeExpanded = false
	__Product__00000018_.IsConsumersNodeExpanded = false
	__Product__00000018_.LayoutDirection = models.Vertical

	__Product__00000019_.Name = `Thermal Sensors`
	__Product__00000019_.ComputedPrefix = `1.1.1.3`
	__Product__00000019_.IsExpanded = false
	__Product__00000019_.IsImport = false
	__Product__00000019_.Description = ``
	__Product__00000019_.IsProducersNodeExpanded = false
	__Product__00000019_.IsConsumersNodeExpanded = false
	__Product__00000019_.LayoutDirection = models.Vertical

	__Product__00000020_.Name = `Dog House`
	__Product__00000020_.ComputedPrefix = `1.1.1.1.1.1.1`
	__Product__00000020_.IsExpanded = false
	__Product__00000020_.IsImport = false
	__Product__00000020_.Description = ``
	__Product__00000020_.IsProducersNodeExpanded = false
	__Product__00000020_.IsConsumersNodeExpanded = false
	__Product__00000020_.LayoutDirection = models.Vertical

	__Product__00000021_.Name = `Orbital Maneuvering and Attitude Control (OMAC), 30 000 lbf`
	__Product__00000021_.ComputedPrefix = `1.1.1.4`
	__Product__00000021_.IsExpanded = false
	__Product__00000021_.IsImport = false
	__Product__00000021_.Description = ``
	__Product__00000021_.IsProducersNodeExpanded = false
	__Product__00000021_.IsConsumersNodeExpanded = false
	__Product__00000021_.LayoutDirection = models.Vertical

	__Product__00000022_.Name = `Boeing’s Enterprise Root Cause/Corrective Actions (eRCCA)`
	__Product__00000022_.ComputedPrefix = `2.3`
	__Product__00000022_.IsExpanded = false
	__Product__00000022_.IsImport = false
	__Product__00000022_.Description = ``
	__Product__00000022_.IsProducersNodeExpanded = false
	__Product__00000022_.IsConsumersNodeExpanded = false
	__Product__00000022_.LayoutDirection = models.Vertical

	__Product__00000023_.Name = `3.1   Orbital Flight Test (OFT) Summary`
	__Product__00000023_.ComputedPrefix = `2.2.2.1`
	__Product__00000023_.IsExpanded = false
	__Product__00000023_.IsImport = false
	__Product__00000023_.Description = ``
	__Product__00000023_.IsProducersNodeExpanded = false
	__Product__00000023_.IsConsumersNodeExpanded = false
	__Product__00000023_.LayoutDirection = models.Vertical

	__Product__00000024_.Name = `3.2   Orbital Flight Test 2 (OFT-2) Summary`
	__Product__00000024_.ComputedPrefix = `2.2.2.2`
	__Product__00000024_.IsExpanded = false
	__Product__00000024_.IsImport = false
	__Product__00000024_.Description = ``
	__Product__00000024_.IsProducersNodeExpanded = false
	__Product__00000024_.IsConsumersNodeExpanded = false
	__Product__00000024_.LayoutDirection = models.Vertical

	__Product__00000025_.Name = `Flight Software (FSW)`
	__Product__00000025_.ComputedPrefix = `1.1.1.5`
	__Product__00000025_.IsExpanded = false
	__Product__00000025_.IsImport = false
	__Product__00000025_.Description = ``
	__Product__00000025_.IsProducersNodeExpanded = false
	__Product__00000025_.IsConsumersNodeExpanded = false
	__Product__00000025_.LayoutDirection = models.Vertical

	__Product__00000026_.Name = `Fault Detection, Isolation, and Recovery (FDIR)`
	__Product__00000026_.ComputedPrefix = `1.1.1.6`
	__Product__00000026_.IsExpanded = false
	__Product__00000026_.IsImport = false
	__Product__00000026_.Description = ``
	__Product__00000026_.IsProducersNodeExpanded = false
	__Product__00000026_.IsConsumersNodeExpanded = false
	__Product__00000026_.LayoutDirection = models.Vertical

	__Product__00000027_.Name = `3.3 Comparing SM RCS Thrusters Triggering Fail-Off FDIR on OFT1/OFT2`
	__Product__00000027_.ComputedPrefix = `2.2.2.3`
	__Product__00000027_.IsExpanded = false
	__Product__00000027_.IsImport = false
	__Product__00000027_.Description = ``
	__Product__00000027_.IsProducersNodeExpanded = false
	__Product__00000027_.IsConsumersNodeExpanded = false
	__Product__00000027_.LayoutDirection = models.Vertical

	__Product__00000028_.Name = `Transducer / pressure sensor (Pc)`
	__Product__00000028_.ComputedPrefix = `1.1.1.1.1.1.2`
	__Product__00000028_.IsExpanded = false
	__Product__00000028_.IsImport = false
	__Product__00000028_.Description = ``
	__Product__00000028_.IsProducersNodeExpanded = false
	__Product__00000028_.IsConsumersNodeExpanded = false
	__Product__00000028_.LayoutDirection = models.Vertical

	__Product__00000029_.Name = `4     Technical Root Cause Analysis (RCA) and Findings`
	__Product__00000029_.ComputedPrefix = `2.2.3`
	__Product__00000029_.IsExpanded = false
	__Product__00000029_.IsImport = false
	__Product__00000029_.Description = ``
	__Product__00000029_.IsProducersNodeExpanded = false
	__Product__00000029_.IsConsumersNodeExpanded = false
	__Product__00000029_.LayoutDirection = models.Vertical

	__Product__00000030_.Name = `4.1   Objectives and Approach `
	__Product__00000030_.ComputedPrefix = `2.2.3.1`
	__Product__00000030_.IsExpanded = false
	__Product__00000030_.IsImport = false
	__Product__00000030_.Description = ``
	__Product__00000030_.IsProducersNodeExpanded = false
	__Product__00000030_.IsConsumersNodeExpanded = false
	__Product__00000030_.LayoutDirection = models.Vertical

	__Product__00000031_.Name = `4.2 Definitions`
	__Product__00000031_.ComputedPrefix = `2.2.3.2`
	__Product__00000031_.IsExpanded = false
	__Product__00000031_.IsImport = false
	__Product__00000031_.Description = ``
	__Product__00000031_.IsProducersNodeExpanded = false
	__Product__00000031_.IsConsumersNodeExpanded = false
	__Product__00000031_.LayoutDirection = models.Vertical

	__Product__00000032_.Name = `4.3 Fault Tree`
	__Product__00000032_.ComputedPrefix = `2.2.3.3`
	__Product__00000032_.IsExpanded = false
	__Product__00000032_.IsImport = false
	__Product__00000032_.Description = ``
	__Product__00000032_.IsProducersNodeExpanded = false
	__Product__00000032_.IsConsumersNodeExpanded = false
	__Product__00000032_.LayoutDirection = models.Vertical

	__Product__00000033_.Name = `Guidance, Navigation, and Control (GNC)`
	__Product__00000033_.ComputedPrefix = `1.1.1.7`
	__Product__00000033_.IsExpanded = false
	__Product__00000033_.IsImport = false
	__Product__00000033_.Description = ``
	__Product__00000033_.IsProducersNodeExpanded = false
	__Product__00000033_.IsConsumersNodeExpanded = false
	__Product__00000033_.LayoutDirection = models.Vertical

	__Product__00000034_.Name = `4.4   Analysis: CM RCS Jet Failure`
	__Product__00000034_.ComputedPrefix = `2.2.3.4`
	__Product__00000034_.IsExpanded = false
	__Product__00000034_.IsImport = false
	__Product__00000034_.Description = ``
	__Product__00000034_.IsProducersNodeExpanded = false
	__Product__00000034_.IsConsumersNodeExpanded = false
	__Product__00000034_.LayoutDirection = models.Vertical

	__Product__00000035_.Name = `4.5   Analysis: Loss of 6DOF Control - SM RCS Jet Failures`
	__Product__00000035_.ComputedPrefix = `2.2.3.5`
	__Product__00000035_.IsExpanded = false
	__Product__00000035_.IsImport = false
	__Product__00000035_.Description = ``
	__Product__00000035_.IsProducersNodeExpanded = false
	__Product__00000035_.IsConsumersNodeExpanded = false
	__Product__00000035_.LayoutDirection = models.Vertical

	__Product__00000036_.Name = `4.5.1 Description of the system`
	__Product__00000036_.ComputedPrefix = `2.2.3.5.1`
	__Product__00000036_.IsExpanded = false
	__Product__00000036_.IsImport = false
	__Product__00000036_.Description = ``
	__Product__00000036_.IsProducersNodeExpanded = false
	__Product__00000036_.IsConsumersNodeExpanded = false
	__Product__00000036_.LayoutDirection = models.Vertical

	__Product__00000037_.Name = ``
	__Product__00000037_.ComputedPrefix = `2.2.3.6`
	__Product__00000037_.IsExpanded = false
	__Product__00000037_.IsImport = false
	__Product__00000037_.Description = ``
	__Product__00000037_.IsProducersNodeExpanded = false
	__Product__00000037_.IsConsumersNodeExpanded = false
	__Product__00000037_.LayoutDirection = models.Vertical

	__Product__00000038_.Name = `Inertial Measurement Unit (IMU) `
	__Product__00000038_.ComputedPrefix = `1.1.1.7.1`
	__Product__00000038_.IsExpanded = false
	__Product__00000038_.IsImport = false
	__Product__00000038_.Description = ``
	__Product__00000038_.IsProducersNodeExpanded = false
	__Product__00000038_.IsConsumersNodeExpanded = false
	__Product__00000038_.LayoutDirection = models.Vertical

	__Product__00000039_.Name = `4.5.2   Description of Events and Timeline `
	__Product__00000039_.ComputedPrefix = `2.2.3.5.2`
	__Product__00000039_.IsExpanded = false
	__Product__00000039_.IsImport = false
	__Product__00000039_.Description = ``
	__Product__00000039_.IsProducersNodeExpanded = false
	__Product__00000039_.IsConsumersNodeExpanded = false
	__Product__00000039_.LayoutDirection = models.Vertical

	__Product__00000040_.Name = `B1A3 Thruster failure`
	__Product__00000040_.ComputedPrefix = `2.2.3.5.3`
	__Product__00000040_.IsExpanded = false
	__Product__00000040_.IsImport = false
	__Product__00000040_.Description = ``
	__Product__00000040_.IsProducersNodeExpanded = false
	__Product__00000040_.IsConsumersNodeExpanded = false
	__Product__00000040_.LayoutDirection = models.Vertical

	__Product__00000041_.Name = `4.5.4   Starliner Engine Testing at White Sands Test Facility (WSTF) during CFT`
	__Product__00000041_.ComputedPrefix = `2.2.3.5.4`
	__Product__00000041_.IsExpanded = false
	__Product__00000041_.IsImport = false
	__Product__00000041_.Description = ``
	__Product__00000041_.IsProducersNodeExpanded = false
	__Product__00000041_.IsConsumersNodeExpanded = false
	__Product__00000041_.LayoutDirection = models.Vertical

	__Product__00000042_.Name = `4.5.5 Fault Tree`
	__Product__00000042_.ComputedPrefix = `2.2.3.5.5`
	__Product__00000042_.IsExpanded = false
	__Product__00000042_.IsImport = false
	__Product__00000042_.Description = ``
	__Product__00000042_.IsProducersNodeExpanded = false
	__Product__00000042_.IsConsumersNodeExpanded = false
	__Product__00000042_.LayoutDirection = models.Vertical

	__Product__00000043_.Name = `4.5.6 Most Probable Proximate Cause`
	__Product__00000043_.ComputedPrefix = `2.2.3.5.6`
	__Product__00000043_.IsExpanded = false
	__Product__00000043_.IsImport = false
	__Product__00000043_.Description = ``
	__Product__00000043_.IsProducersNodeExpanded = false
	__Product__00000043_.IsConsumersNodeExpanded = false
	__Product__00000043_.LayoutDirection = models.Vertical

	__Product__00000044_.Name = `Intermediate Cause 5: Inadequate Thruster Thermal Models Inadequate thermal modelling caused insufficient scrutiny for the thermal environment, leading to excessive heating from RCS thermal soakback and integrated heating from OMACs.`
	__Product__00000044_.ComputedPrefix = `2.2.3.5.6.1`
	__Product__00000044_.IsExpanded = false
	__Product__00000044_.IsImport = false
	__Product__00000044_.Description = ``
	__Product__00000044_.IsProducersNodeExpanded = false
	__Product__00000044_.IsConsumersNodeExpanded = false
	__Product__00000044_.LayoutDirection = models.Vertical

	__Product__00000045_.Name = `4.6 Analysis: Helium Leak`
	__Product__00000045_.ComputedPrefix = `2.2.3.5.7`
	__Product__00000045_.IsExpanded = false
	__Product__00000045_.IsImport = false
	__Product__00000045_.Description = ``
	__Product__00000045_.IsProducersNodeExpanded = false
	__Product__00000045_.IsConsumersNodeExpanded = false
	__Product__00000045_.LayoutDirection = models.Vertical

	__Product__00000046_.Name = `4.6.1 Description of the system`
	__Product__00000046_.ComputedPrefix = `2.2.3.5.7.1`
	__Product__00000046_.IsExpanded = false
	__Product__00000046_.IsImport = false
	__Product__00000046_.Description = ``
	__Product__00000046_.IsProducersNodeExpanded = false
	__Product__00000046_.IsConsumersNodeExpanded = false
	__Product__00000046_.LayoutDirection = models.Vertical

	__ProductCompositionShape__00000000_.Name = `Reports to `
	__ProductCompositionShape__00000000_.StartRatio = 0.500000
	__ProductCompositionShape__00000000_.EndRatio = 0.500000
	__ProductCompositionShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000000_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000000_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000000_.IsHidden = false

	__ProductCompositionShape__00000002_.Name = `Commercial Crew Transportation Capability (CCtCap). to Starliner`
	__ProductCompositionShape__00000002_.StartRatio = 0.500000
	__ProductCompositionShape__00000002_.EndRatio = 0.500000
	__ProductCompositionShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000002_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000002_.IsHidden = false

	__ProductCompositionShape__00000003_.Name = `Commercial Crew Transportation Capability (CCtCap). to Dragon`
	__ProductCompositionShape__00000003_.StartRatio = 0.500000
	__ProductCompositionShape__00000003_.EndRatio = 0.500000
	__ProductCompositionShape__00000003_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000003_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000003_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000003_.IsHidden = false

	__ProductCompositionShape__00000004_.Name = `NASA Assets/Capabities to Commercial Crew Transportation Capability (CCtCap).`
	__ProductCompositionShape__00000004_.StartRatio = 0.500000
	__ProductCompositionShape__00000004_.EndRatio = 0.500000
	__ProductCompositionShape__00000004_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000004_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000004_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000004_.IsHidden = false

	__ProductCompositionShape__00000005_.Name = `Commercial Crew Transportation Capability (CCtCap). to `
	__ProductCompositionShape__00000005_.StartRatio = 0.501832
	__ProductCompositionShape__00000005_.EndRatio = 0.300650
	__ProductCompositionShape__00000005_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000005_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000005_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000005_.IsHidden = false

	__ProductCompositionShape__00000006_.Name = `Program Investigation Team (PIT) Report to `
	__ProductCompositionShape__00000006_.StartRatio = 0.784088
	__ProductCompositionShape__00000006_.EndRatio = 0.460874
	__ProductCompositionShape__00000006_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000006_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000006_.CornerOffsetRatio = -0.073041
	__ProductCompositionShape__00000006_.IsHidden = false

	__ProductCompositionShape__00000007_.Name = `CM to `
	__ProductCompositionShape__00000007_.StartRatio = 0.500000
	__ProductCompositionShape__00000007_.EndRatio = 0.500000
	__ProductCompositionShape__00000007_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000007_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000007_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000007_.IsHidden = false

	__ProductCompositionShape__00000008_.Name = `RCS to `
	__ProductCompositionShape__00000008_.StartRatio = 0.500000
	__ProductCompositionShape__00000008_.EndRatio = 0.500000
	__ProductCompositionShape__00000008_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000008_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000008_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000008_.IsHidden = false

	__ProductCompositionShape__00000009_.Name = `CST-100 Starliner to CM`
	__ProductCompositionShape__00000009_.StartRatio = 0.500000
	__ProductCompositionShape__00000009_.EndRatio = 0.500000
	__ProductCompositionShape__00000009_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000009_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000009_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000009_.IsHidden = false

	__ProductCompositionShape__00000010_.Name = `CST-100 Starliner to `
	__ProductCompositionShape__00000010_.StartRatio = 0.500000
	__ProductCompositionShape__00000010_.EndRatio = 0.500000
	__ProductCompositionShape__00000010_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000010_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000010_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000010_.IsHidden = false

	__ProductCompositionShape__00000011_.Name = `Service Module (SM) to `
	__ProductCompositionShape__00000011_.StartRatio = 0.500000
	__ProductCompositionShape__00000011_.EndRatio = 0.500000
	__ProductCompositionShape__00000011_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000011_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000011_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000011_.IsHidden = false

	__ProductCompositionShape__00000012_.Name = `SM Reaction Control System (SM RCS) to `
	__ProductCompositionShape__00000012_.StartRatio = 0.500000
	__ProductCompositionShape__00000012_.EndRatio = 0.500000
	__ProductCompositionShape__00000012_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000012_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000012_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000012_.IsHidden = false

	__ProductCompositionShape__00000013_.Name = `Program Investigation Team (PIT) Report to `
	__ProductCompositionShape__00000013_.StartRatio = 0.555517
	__ProductCompositionShape__00000013_.EndRatio = 0.498374
	__ProductCompositionShape__00000013_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000013_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000013_.CornerOffsetRatio = -0.077041
	__ProductCompositionShape__00000013_.IsHidden = false

	__ProductCompositionShape__00000014_.Name = `Chap 2 Additional Investigations  to `
	__ProductCompositionShape__00000014_.StartRatio = 0.500000
	__ProductCompositionShape__00000014_.EndRatio = 0.500000
	__ProductCompositionShape__00000014_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000014_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000014_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000014_.IsHidden = false

	__ProductCompositionShape__00000016_.Name = `12 CM RCS Thrusters to `
	__ProductCompositionShape__00000016_.StartRatio = 0.500000
	__ProductCompositionShape__00000016_.EndRatio = 0.500000
	__ProductCompositionShape__00000016_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000016_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000016_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000016_.IsHidden = false

	__ProductCompositionShape__00000018_.Name = `Reports to Boeing’s Enterprise Root Cause/Corrective Actions (eRCCA)`
	__ProductCompositionShape__00000018_.StartRatio = 0.500000
	__ProductCompositionShape__00000018_.EndRatio = 0.445736
	__ProductCompositionShape__00000018_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000018_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000018_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000018_.IsHidden = false

	__ProductCompositionShape__00000019_.Name = `Chap3 Commercial Crew Program (CCP) Background to `
	__ProductCompositionShape__00000019_.StartRatio = 0.500000
	__ProductCompositionShape__00000019_.EndRatio = 0.500000
	__ProductCompositionShape__00000019_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000019_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000019_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000019_.IsHidden = false

	__ProductCompositionShape__00000020_.Name = `Chap3 Commercial Crew Program (CCP) Background to `
	__ProductCompositionShape__00000020_.StartRatio = 0.500000
	__ProductCompositionShape__00000020_.EndRatio = 0.500000
	__ProductCompositionShape__00000020_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000020_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000020_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000020_.IsHidden = false

	__ProductCompositionShape__00000021_.Name = `CST-100 Starliner to `
	__ProductCompositionShape__00000021_.StartRatio = 0.520485
	__ProductCompositionShape__00000021_.EndRatio = 0.363342
	__ProductCompositionShape__00000021_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000021_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000021_.CornerOffsetRatio = 5.740012
	__ProductCompositionShape__00000021_.IsHidden = false

	__ProductCompositionShape__00000022_.Name = `CST-100 Starliner to `
	__ProductCompositionShape__00000022_.StartRatio = 0.506199
	__ProductCompositionShape__00000022_.EndRatio = 0.149056
	__ProductCompositionShape__00000022_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000022_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000022_.CornerOffsetRatio = 5.740012
	__ProductCompositionShape__00000022_.IsHidden = false

	__ProductCompositionShape__00000023_.Name = `Chap3 Commercial Crew Program (CCP) Background to `
	__ProductCompositionShape__00000023_.StartRatio = 0.500000
	__ProductCompositionShape__00000023_.EndRatio = 0.500000
	__ProductCompositionShape__00000023_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000023_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000023_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000023_.IsHidden = false

	__ProductCompositionShape__00000024_.Name = `12 CM RCS Thrusters (1200 lbf) to `
	__ProductCompositionShape__00000024_.StartRatio = 0.500000
	__ProductCompositionShape__00000024_.EndRatio = 0.500000
	__ProductCompositionShape__00000024_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000024_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000024_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000024_.IsHidden = false

	__ProductCompositionShape__00000025_.Name = `Chap3 Commercial Crew Program (CCP) Background to `
	__ProductCompositionShape__00000025_.StartRatio = 0.855517
	__ProductCompositionShape__00000025_.EndRatio = 0.526945
	__ProductCompositionShape__00000025_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000025_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000025_.CornerOffsetRatio = -0.069041
	__ProductCompositionShape__00000025_.IsHidden = false

	__ProductCompositionShape__00000026_.Name = `4     Technical Root Cause Analysis (RCA) and Findings to `
	__ProductCompositionShape__00000026_.StartRatio = 0.500000
	__ProductCompositionShape__00000026_.EndRatio = 0.500000
	__ProductCompositionShape__00000026_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000026_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000026_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000026_.IsHidden = false

	__ProductCompositionShape__00000027_.Name = `4     Technical Root Cause Analysis (RCA) and Findings to `
	__ProductCompositionShape__00000027_.StartRatio = 0.500000
	__ProductCompositionShape__00000027_.EndRatio = 0.500000
	__ProductCompositionShape__00000027_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000027_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000027_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000027_.IsHidden = false

	__ProductCompositionShape__00000028_.Name = `4     Technical Root Cause Analysis (RCA) and Findings to `
	__ProductCompositionShape__00000028_.StartRatio = 0.500000
	__ProductCompositionShape__00000028_.EndRatio = 0.500000
	__ProductCompositionShape__00000028_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000028_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000028_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000028_.IsHidden = false

	__ProductCompositionShape__00000029_.Name = `CST-100 Starliner to `
	__ProductCompositionShape__00000029_.StartRatio = 0.463243
	__ProductCompositionShape__00000029_.EndRatio = 0.263243
	__ProductCompositionShape__00000029_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000029_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000029_.CornerOffsetRatio = 5.757367
	__ProductCompositionShape__00000029_.IsHidden = false

	__ProductCompositionShape__00000030_.Name = `4     Technical Root Cause Analysis (RCA) and Findings to `
	__ProductCompositionShape__00000030_.StartRatio = 0.500000
	__ProductCompositionShape__00000030_.EndRatio = 0.500000
	__ProductCompositionShape__00000030_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000030_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000030_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000030_.IsHidden = false

	__ProductCompositionShape__00000031_.Name = `4     Technical Root Cause Analysis (RCA) and Findings to `
	__ProductCompositionShape__00000031_.StartRatio = 0.826945
	__ProductCompositionShape__00000031_.EndRatio = 0.641231
	__ProductCompositionShape__00000031_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000031_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000031_.CornerOffsetRatio = -0.117041
	__ProductCompositionShape__00000031_.IsHidden = false

	__ProductCompositionShape__00000032_.Name = `4.5   Analysis: Loss of 6DOF Control - SM RCS Jet Failures to `
	__ProductCompositionShape__00000032_.StartRatio = 0.500000
	__ProductCompositionShape__00000032_.EndRatio = 0.500000
	__ProductCompositionShape__00000032_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000032_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000032_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000032_.IsHidden = false

	__ProductCompositionShape__00000033_.Name = `4     Technical Root Cause Analysis (RCA) and Findings to `
	__ProductCompositionShape__00000033_.StartRatio = 0.500000
	__ProductCompositionShape__00000033_.EndRatio = 0.500000
	__ProductCompositionShape__00000033_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000033_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000033_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000033_.IsHidden = false

	__ProductCompositionShape__00000034_.Name = `Guidance, Navigation, and Control (GNC) to `
	__ProductCompositionShape__00000034_.StartRatio = 0.500000
	__ProductCompositionShape__00000034_.EndRatio = 0.500000
	__ProductCompositionShape__00000034_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000034_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000034_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000034_.IsHidden = false

	__ProductCompositionShape__00000035_.Name = `CST-100 Starliner to Orbital Maneuvering and Attitude Control (OMAC), 30 000 lbf`
	__ProductCompositionShape__00000035_.StartRatio = 0.500000
	__ProductCompositionShape__00000035_.EndRatio = 0.500000
	__ProductCompositionShape__00000035_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000035_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000035_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000035_.IsHidden = false

	__ProductCompositionShape__00000036_.Name = `4.5   Analysis: Loss of 6DOF Control - SM RCS Jet Failures to `
	__ProductCompositionShape__00000036_.StartRatio = 0.500000
	__ProductCompositionShape__00000036_.EndRatio = 0.500000
	__ProductCompositionShape__00000036_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000036_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000036_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000036_.IsHidden = false

	__ProductCompositionShape__00000037_.Name = `4.5   Analysis: Loss of 6DOF Control - SM RCS Jet Failures to `
	__ProductCompositionShape__00000037_.StartRatio = 0.500000
	__ProductCompositionShape__00000037_.EndRatio = 0.500000
	__ProductCompositionShape__00000037_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000037_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000037_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000037_.IsHidden = false

	__ProductCompositionShape__00000038_.Name = `4.5   Analysis: Loss of 6DOF Control - SM RCS Jet Failures to `
	__ProductCompositionShape__00000038_.StartRatio = 0.500000
	__ProductCompositionShape__00000038_.EndRatio = 0.500000
	__ProductCompositionShape__00000038_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000038_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000038_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000038_.IsHidden = false

	__ProductCompositionShape__00000039_.Name = `4.5   Analysis: Loss of 6DOF Control - SM RCS Jet Failures to `
	__ProductCompositionShape__00000039_.StartRatio = 0.500000
	__ProductCompositionShape__00000039_.EndRatio = 0.500000
	__ProductCompositionShape__00000039_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000039_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000039_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000039_.IsHidden = false

	__ProductCompositionShape__00000040_.Name = `4.5   Analysis: Loss of 6DOF Control - SM RCS Jet Failures to `
	__ProductCompositionShape__00000040_.StartRatio = 0.914445
	__ProductCompositionShape__00000040_.EndRatio = 0.471588
	__ProductCompositionShape__00000040_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000040_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000040_.CornerOffsetRatio = -0.139307
	__ProductCompositionShape__00000040_.IsHidden = false

	__ProductCompositionShape__00000041_.Name = `4.5.6 Most Probable Proximate Cause to `
	__ProductCompositionShape__00000041_.StartRatio = 0.500000
	__ProductCompositionShape__00000041_.EndRatio = 0.500000
	__ProductCompositionShape__00000041_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000041_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000041_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000041_.IsHidden = false

	__ProductCompositionShape__00000042_.Name = `4.5   Analysis: Loss of 6DOF Control - SM RCS Jet Failures to `
	__ProductCompositionShape__00000042_.StartRatio = 0.828731
	__ProductCompositionShape__00000042_.EndRatio = 0.385874
	__ProductCompositionShape__00000042_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000042_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000042_.CornerOffsetRatio = -0.165041
	__ProductCompositionShape__00000042_.IsHidden = false

	__ProductCompositionShape__00000043_.Name = `4.6 Analysis: Helium Leak to `
	__ProductCompositionShape__00000043_.StartRatio = 0.500000
	__ProductCompositionShape__00000043_.EndRatio = 0.500000
	__ProductCompositionShape__00000043_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000043_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000043_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000043_.IsHidden = false

	__ProductCompositionShape__00000044_.Name = `Reports to Program Investigation Team (PIT) Reports`
	__ProductCompositionShape__00000044_.StartRatio = 0.500000
	__ProductCompositionShape__00000044_.EndRatio = 0.500000
	__ProductCompositionShape__00000044_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000044_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000044_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000044_.IsHidden = false

	__ProductShape__00000002_.Name = `Dragon-PBS`
	__ProductShape__00000002_.OverideLayoutDirection = false
	__ProductShape__00000002_.LayoutDirection = models.Vertical
	__ProductShape__00000002_.X = 350.000000
	__ProductShape__00000002_.Y = 330.000000
	__ProductShape__00000002_.Width = 250.000000
	__ProductShape__00000002_.Height = 70.000000
	__ProductShape__00000002_.IsHidden = false

	__ProductShape__00000003_.Name = `Starliner-PBS`
	__ProductShape__00000003_.OverideLayoutDirection = false
	__ProductShape__00000003_.LayoutDirection = models.Vertical
	__ProductShape__00000003_.X = 50.000000
	__ProductShape__00000003_.Y = 330.000000
	__ProductShape__00000003_.Width = 250.000000
	__ProductShape__00000003_.Height = 70.000000
	__ProductShape__00000003_.IsHidden = false

	__ProductShape__00000004_.Name = `-PBS`
	__ProductShape__00000004_.OverideLayoutDirection = false
	__ProductShape__00000004_.LayoutDirection = models.Vertical
	__ProductShape__00000004_.X = 950.000000
	__ProductShape__00000004_.Y = 50.000000
	__ProductShape__00000004_.Width = 250.000000
	__ProductShape__00000004_.Height = 70.000000
	__ProductShape__00000004_.IsHidden = false

	__ProductShape__00000005_.Name = `-PBS`
	__ProductShape__00000005_.OverideLayoutDirection = false
	__ProductShape__00000005_.LayoutDirection = models.Vertical
	__ProductShape__00000005_.X = 1125.000000
	__ProductShape__00000005_.Y = 135.000000
	__ProductShape__00000005_.Width = 250.000000
	__ProductShape__00000005_.Height = 70.000000
	__ProductShape__00000005_.IsHidden = false

	__ProductShape__00000006_.Name = `-PBS`
	__ProductShape__00000006_.OverideLayoutDirection = false
	__ProductShape__00000006_.LayoutDirection = models.Vertical
	__ProductShape__00000006_.X = 1125.000000
	__ProductShape__00000006_.Y = 220.000000
	__ProductShape__00000006_.Width = 250.000000
	__ProductShape__00000006_.Height = 70.000000
	__ProductShape__00000006_.IsHidden = false

	__ProductShape__00000007_.Name = `Program Investigation Team (PIT) Report-PIT focus`
	__ProductShape__00000007_.OverideLayoutDirection = false
	__ProductShape__00000007_.LayoutDirection = models.Vertical
	__ProductShape__00000007_.X = 912.017987
	__ProductShape__00000007_.Y = 275.867255
	__ProductShape__00000007_.Width = 250.000000
	__ProductShape__00000007_.Height = 70.000000
	__ProductShape__00000007_.IsHidden = false

	__ProductShape__00000008_.Name = ` Starliner Tests and Anomalies Review (STAR) Investigation Report-PIT focus`
	__ProductShape__00000008_.OverideLayoutDirection = false
	__ProductShape__00000008_.LayoutDirection = models.Vertical
	__ProductShape__00000008_.X = 452.865680
	__ProductShape__00000008_.Y = 64.093982
	__ProductShape__00000008_.Width = 250.000000
	__ProductShape__00000008_.Height = 70.000000
	__ProductShape__00000008_.IsHidden = false

	__ProductShape__00000009_.Name = `-PBS`
	__ProductShape__00000009_.OverideLayoutDirection = false
	__ProductShape__00000009_.LayoutDirection = models.Vertical
	__ProductShape__00000009_.X = 50.000000
	__ProductShape__00000009_.Y = 190.000000
	__ProductShape__00000009_.Width = 250.000000
	__ProductShape__00000009_.Height = 70.000000
	__ProductShape__00000009_.IsHidden = false

	__ProductShape__00000010_.Name = `-PBS`
	__ProductShape__00000010_.OverideLayoutDirection = false
	__ProductShape__00000010_.LayoutDirection = models.Vertical
	__ProductShape__00000010_.X = 50.000000
	__ProductShape__00000010_.Y = 50.000000
	__ProductShape__00000010_.Width = 250.000000
	__ProductShape__00000010_.Height = 70.000000
	__ProductShape__00000010_.IsHidden = false

	__ProductShape__00000011_.Name = `-PBS`
	__ProductShape__00000011_.OverideLayoutDirection = false
	__ProductShape__00000011_.LayoutDirection = models.Vertical
	__ProductShape__00000011_.X = 650.000000
	__ProductShape__00000011_.Y = 190.000000
	__ProductShape__00000011_.Width = 250.000000
	__ProductShape__00000011_.Height = 70.000000
	__ProductShape__00000011_.IsHidden = false

	__ProductShape__00000012_.Name = `Program Investigation Team (PIT) Report-PIT Report`
	__ProductShape__00000012_.OverideLayoutDirection = false
	__ProductShape__00000012_.LayoutDirection = models.Vertical
	__ProductShape__00000012_.X = 80.516205
	__ProductShape__00000012_.Y = 109.119696
	__ProductShape__00000012_.Width = 250.000000
	__ProductShape__00000012_.Height = 70.000000
	__ProductShape__00000012_.IsHidden = false

	__ProductShape__00000013_.Name = `-PIT Report`
	__ProductShape__00000013_.OverideLayoutDirection = false
	__ProductShape__00000013_.LayoutDirection = models.Vertical
	__ProductShape__00000013_.X = 82.516205
	__ProductShape__00000013_.Y = 559.119696
	__ProductShape__00000013_.Width = 276.000000
	__ProductShape__00000013_.Height = 70.000000
	__ProductShape__00000013_.IsHidden = false

	__ProductShape__00000014_.Name = `-PIT Report`
	__ProductShape__00000014_.OverideLayoutDirection = false
	__ProductShape__00000014_.LayoutDirection = models.Vertical
	__ProductShape__00000014_.X = 766.767652
	__ProductShape__00000014_.Y = 110.422064
	__ProductShape__00000014_.Width = 250.000000
	__ProductShape__00000014_.Height = 70.000000
	__ProductShape__00000014_.IsHidden = false

	__ProductShape__00000015_.Name = `-RCS PBS`
	__ProductShape__00000015_.OverideLayoutDirection = false
	__ProductShape__00000015_.LayoutDirection = models.Vertical
	__ProductShape__00000015_.X = 34.373886
	__ProductShape__00000015_.Y = 167.792999
	__ProductShape__00000015_.Width = 250.000000
	__ProductShape__00000015_.Height = 70.000000
	__ProductShape__00000015_.IsHidden = false

	__ProductShape__00000016_.Name = `-RCS PBS`
	__ProductShape__00000016_.OverideLayoutDirection = false
	__ProductShape__00000016_.LayoutDirection = models.Vertical
	__ProductShape__00000016_.X = 34.373886
	__ProductShape__00000016_.Y = 307.792999
	__ProductShape__00000016_.Width = 250.000000
	__ProductShape__00000016_.Height = 70.000000
	__ProductShape__00000016_.IsHidden = false

	__ProductShape__00000017_.Name = `-RCS PBS`
	__ProductShape__00000017_.OverideLayoutDirection = false
	__ProductShape__00000017_.LayoutDirection = models.Vertical
	__ProductShape__00000017_.X = 34.373886
	__ProductShape__00000017_.Y = 447.792999
	__ProductShape__00000017_.Width = 250.000000
	__ProductShape__00000017_.Height = 70.000000
	__ProductShape__00000017_.IsHidden = false

	__ProductShape__00000018_.Name = `CST-100 Starliner-RCS PBS`
	__ProductShape__00000018_.OverideLayoutDirection = false
	__ProductShape__00000018_.LayoutDirection = models.Vertical
	__ProductShape__00000018_.X = 33.736439
	__ProductShape__00000018_.Y = 23.572991
	__ProductShape__00000018_.Width = 250.000000
	__ProductShape__00000018_.Height = 70.000000
	__ProductShape__00000018_.IsHidden = false

	__ProductShape__00000019_.Name = `-RCS PBS`
	__ProductShape__00000019_.OverideLayoutDirection = false
	__ProductShape__00000019_.LayoutDirection = models.Vertical
	__ProductShape__00000019_.X = 595.736439
	__ProductShape__00000019_.Y = 182.573022
	__ProductShape__00000019_.Width = 250.000000
	__ProductShape__00000019_.Height = 70.000000
	__ProductShape__00000019_.IsHidden = false

	__ProductShape__00000020_.Name = `-RCS PBS`
	__ProductShape__00000020_.OverideLayoutDirection = false
	__ProductShape__00000020_.LayoutDirection = models.Vertical
	__ProductShape__00000020_.X = 595.736439
	__ProductShape__00000020_.Y = 322.573022
	__ProductShape__00000020_.Width = 250.000000
	__ProductShape__00000020_.Height = 70.000000
	__ProductShape__00000020_.IsHidden = false

	__ProductShape__00000021_.Name = `-RCS PBS`
	__ProductShape__00000021_.OverideLayoutDirection = false
	__ProductShape__00000021_.LayoutDirection = models.Vertical
	__ProductShape__00000021_.X = 595.736439
	__ProductShape__00000021_.Y = 462.573022
	__ProductShape__00000021_.Width = 250.000000
	__ProductShape__00000021_.Height = 70.000000
	__ProductShape__00000021_.IsHidden = false

	__ProductShape__00000022_.Name = `-PIT Report`
	__ProductShape__00000022_.OverideLayoutDirection = false
	__ProductShape__00000022_.LayoutDirection = models.Vertical
	__ProductShape__00000022_.X = 99.516205
	__ProductShape__00000022_.Y = 315.119696
	__ProductShape__00000022_.Width = 250.000000
	__ProductShape__00000022_.Height = 70.000000
	__ProductShape__00000022_.IsHidden = false

	__ProductShape__00000023_.Name = `-PIT Report`
	__ProductShape__00000023_.OverideLayoutDirection = false
	__ProductShape__00000023_.LayoutDirection = models.Vertical
	__ProductShape__00000023_.X = 232.516205
	__ProductShape__00000023_.Y = 459.119696
	__ProductShape__00000023_.Width = 285.000000
	__ProductShape__00000023_.Height = 70.000000
	__ProductShape__00000023_.IsHidden = false

	__ProductShape__00000024_.Name = `-RCS PBS`
	__ProductShape__00000024_.OverideLayoutDirection = false
	__ProductShape__00000024_.LayoutDirection = models.Vertical
	__ProductShape__00000024_.X = 949.373825
	__ProductShape__00000024_.Y = 158.793044
	__ProductShape__00000024_.Width = 250.000000
	__ProductShape__00000024_.Height = 70.000000
	__ProductShape__00000024_.IsHidden = false

	__ProductShape__00000025_.Name = `-RCS PBS`
	__ProductShape__00000025_.OverideLayoutDirection = false
	__ProductShape__00000025_.LayoutDirection = models.Vertical
	__ProductShape__00000025_.X = 34.373886
	__ProductShape__00000025_.Y = 587.792999
	__ProductShape__00000025_.Width = 250.000000
	__ProductShape__00000025_.Height = 70.000000
	__ProductShape__00000025_.IsHidden = false

	__ProductShape__00000027_.Name = `Boeing’s Enterprise Root Cause/Corrective Actions (eRCCA)-PBS`
	__ProductShape__00000027_.OverideLayoutDirection = false
	__ProductShape__00000027_.LayoutDirection = models.Vertical
	__ProductShape__00000027_.X = 1125.000000
	__ProductShape__00000027_.Y = 305.000000
	__ProductShape__00000027_.Width = 250.000000
	__ProductShape__00000027_.Height = 86.000000
	__ProductShape__00000027_.IsHidden = false

	__ProductShape__00000028_.Name = `-PIT Report`
	__ProductShape__00000028_.OverideLayoutDirection = false
	__ProductShape__00000028_.LayoutDirection = models.Vertical
	__ProductShape__00000028_.X = 223.516205
	__ProductShape__00000028_.Y = 706.119696
	__ProductShape__00000028_.Width = 250.000000
	__ProductShape__00000028_.Height = 70.000000
	__ProductShape__00000028_.IsHidden = false

	__ProductShape__00000029_.Name = `-PIT Report`
	__ProductShape__00000029_.OverideLayoutDirection = false
	__ProductShape__00000029_.LayoutDirection = models.Vertical
	__ProductShape__00000029_.X = 523.516205
	__ProductShape__00000029_.Y = 706.119696
	__ProductShape__00000029_.Width = 250.000000
	__ProductShape__00000029_.Height = 70.000000
	__ProductShape__00000029_.IsHidden = false

	__ProductShape__00000030_.Name = `-RCS PBS`
	__ProductShape__00000030_.OverideLayoutDirection = false
	__ProductShape__00000030_.LayoutDirection = models.Vertical
	__ProductShape__00000030_.X = 1171.736439
	__ProductShape__00000030_.Y = 335.572976
	__ProductShape__00000030_.Width = 250.000000
	__ProductShape__00000030_.Height = 70.000000
	__ProductShape__00000030_.IsHidden = false

	__ProductShape__00000031_.Name = `-RCS PBS`
	__ProductShape__00000031_.OverideLayoutDirection = false
	__ProductShape__00000031_.LayoutDirection = models.Vertical
	__ProductShape__00000031_.X = 1174.736439
	__ProductShape__00000031_.Y = 421.572976
	__ProductShape__00000031_.Width = 250.000000
	__ProductShape__00000031_.Height = 70.000000
	__ProductShape__00000031_.IsHidden = false

	__ProductShape__00000032_.Name = `-PIT Report`
	__ProductShape__00000032_.OverideLayoutDirection = false
	__ProductShape__00000032_.LayoutDirection = models.Vertical
	__ProductShape__00000032_.X = 823.516205
	__ProductShape__00000032_.Y = 706.119696
	__ProductShape__00000032_.Width = 250.000000
	__ProductShape__00000032_.Height = 70.000000
	__ProductShape__00000032_.IsHidden = false

	__ProductShape__00000033_.Name = `-RCS PBS`
	__ProductShape__00000033_.OverideLayoutDirection = false
	__ProductShape__00000033_.LayoutDirection = models.Vertical
	__ProductShape__00000033_.X = 334.373886
	__ProductShape__00000033_.Y = 587.792999
	__ProductShape__00000033_.Width = 250.000000
	__ProductShape__00000033_.Height = 70.000000
	__ProductShape__00000033_.IsHidden = false

	__ProductShape__00000034_.Name = `-PIT Report`
	__ProductShape__00000034_.OverideLayoutDirection = false
	__ProductShape__00000034_.LayoutDirection = models.Vertical
	__ProductShape__00000034_.X = 86.516205
	__ProductShape__00000034_.Y = 819.119696
	__ProductShape__00000034_.Width = 250.000000
	__ProductShape__00000034_.Height = 70.000000
	__ProductShape__00000034_.IsHidden = false

	__ProductShape__00000035_.Name = `-PIT Report`
	__ProductShape__00000035_.OverideLayoutDirection = false
	__ProductShape__00000035_.LayoutDirection = models.Vertical
	__ProductShape__00000035_.X = 86.516205
	__ProductShape__00000035_.Y = 959.119696
	__ProductShape__00000035_.Width = 250.000000
	__ProductShape__00000035_.Height = 70.000000
	__ProductShape__00000035_.IsHidden = false

	__ProductShape__00000036_.Name = `-PIT Report`
	__ProductShape__00000036_.OverideLayoutDirection = false
	__ProductShape__00000036_.LayoutDirection = models.Vertical
	__ProductShape__00000036_.X = 386.516205
	__ProductShape__00000036_.Y = 959.119696
	__ProductShape__00000036_.Width = 250.000000
	__ProductShape__00000036_.Height = 70.000000
	__ProductShape__00000036_.IsHidden = false

	__ProductShape__00000037_.Name = `-PIT Report`
	__ProductShape__00000037_.OverideLayoutDirection = false
	__ProductShape__00000037_.LayoutDirection = models.Vertical
	__ProductShape__00000037_.X = 686.516205
	__ProductShape__00000037_.Y = 959.119696
	__ProductShape__00000037_.Width = 250.000000
	__ProductShape__00000037_.Height = 70.000000
	__ProductShape__00000037_.IsHidden = false

	__ProductShape__00000038_.Name = `-RCS PBS`
	__ProductShape__00000038_.OverideLayoutDirection = false
	__ProductShape__00000038_.LayoutDirection = models.Vertical
	__ProductShape__00000038_.X = 1173.736439
	__ProductShape__00000038_.Y = 509.572991
	__ProductShape__00000038_.Width = 250.000000
	__ProductShape__00000038_.Height = 70.000000
	__ProductShape__00000038_.IsHidden = false

	__ProductShape__00000039_.Name = `-PIT Report`
	__ProductShape__00000039_.OverideLayoutDirection = false
	__ProductShape__00000039_.LayoutDirection = models.Vertical
	__ProductShape__00000039_.X = 986.516205
	__ProductShape__00000039_.Y = 959.119696
	__ProductShape__00000039_.Width = 250.000000
	__ProductShape__00000039_.Height = 70.000000
	__ProductShape__00000039_.IsHidden = false

	__ProductShape__00000040_.Name = `-PIT Report`
	__ProductShape__00000040_.OverideLayoutDirection = false
	__ProductShape__00000040_.LayoutDirection = models.Vertical
	__ProductShape__00000040_.X = 86.516205
	__ProductShape__00000040_.Y = 1096.119696
	__ProductShape__00000040_.Width = 250.000000
	__ProductShape__00000040_.Height = 70.000000
	__ProductShape__00000040_.IsHidden = false

	__ProductShape__00000041_.Name = `-PIT Report`
	__ProductShape__00000041_.OverideLayoutDirection = false
	__ProductShape__00000041_.LayoutDirection = models.Vertical
	__ProductShape__00000041_.X = 83.516205
	__ProductShape__00000041_.Y = 1237.119696
	__ProductShape__00000041_.Width = 250.000000
	__ProductShape__00000041_.Height = 70.000000
	__ProductShape__00000041_.IsHidden = false

	__ProductShape__00000042_.Name = `-PIT Report`
	__ProductShape__00000042_.OverideLayoutDirection = false
	__ProductShape__00000042_.LayoutDirection = models.Vertical
	__ProductShape__00000042_.X = 542.516205
	__ProductShape__00000042_.Y = 1105.119696
	__ProductShape__00000042_.Width = 250.000000
	__ProductShape__00000042_.Height = 70.000000
	__ProductShape__00000042_.IsHidden = false

	__ProductShape__00000043_.Name = `-RCS PBS`
	__ProductShape__00000043_.OverideLayoutDirection = false
	__ProductShape__00000043_.LayoutDirection = models.Vertical
	__ProductShape__00000043_.X = 1173.736439
	__ProductShape__00000043_.Y = 649.572991
	__ProductShape__00000043_.Width = 250.000000
	__ProductShape__00000043_.Height = 70.000000
	__ProductShape__00000043_.IsHidden = false

	__ProductShape__00000044_.Name = `Orbital Maneuvering and Attitude Control (OMAC), 30 000 lbf-RCS PBS`
	__ProductShape__00000044_.OverideLayoutDirection = false
	__ProductShape__00000044_.LayoutDirection = models.Vertical
	__ProductShape__00000044_.X = 300.725524
	__ProductShape__00000044_.Y = 254.217082
	__ProductShape__00000044_.Width = 250.000000
	__ProductShape__00000044_.Height = 70.000000
	__ProductShape__00000044_.IsHidden = false

	__ProductShape__00000045_.Name = `-PIT Report`
	__ProductShape__00000045_.OverideLayoutDirection = false
	__ProductShape__00000045_.LayoutDirection = models.Vertical
	__ProductShape__00000045_.X = 386.516205
	__ProductShape__00000045_.Y = 1236.119696
	__ProductShape__00000045_.Width = 250.000000
	__ProductShape__00000045_.Height = 70.000000
	__ProductShape__00000045_.IsHidden = false

	__ProductShape__00000046_.Name = `-PIT Report`
	__ProductShape__00000046_.OverideLayoutDirection = false
	__ProductShape__00000046_.LayoutDirection = models.Vertical
	__ProductShape__00000046_.X = 686.516205
	__ProductShape__00000046_.Y = 1236.119696
	__ProductShape__00000046_.Width = 250.000000
	__ProductShape__00000046_.Height = 70.000000
	__ProductShape__00000046_.IsHidden = false

	__ProductShape__00000047_.Name = `-PIT Report`
	__ProductShape__00000047_.OverideLayoutDirection = false
	__ProductShape__00000047_.LayoutDirection = models.Vertical
	__ProductShape__00000047_.X = 986.516205
	__ProductShape__00000047_.Y = 1236.119696
	__ProductShape__00000047_.Width = 250.000000
	__ProductShape__00000047_.Height = 70.000000
	__ProductShape__00000047_.IsHidden = false

	__ProductShape__00000048_.Name = `-PIT Report`
	__ProductShape__00000048_.OverideLayoutDirection = false
	__ProductShape__00000048_.LayoutDirection = models.Vertical
	__ProductShape__00000048_.X = 1286.516205
	__ProductShape__00000048_.Y = 1236.119696
	__ProductShape__00000048_.Width = 250.000000
	__ProductShape__00000048_.Height = 70.000000
	__ProductShape__00000048_.IsHidden = false

	__ProductShape__00000049_.Name = `-PIT Report`
	__ProductShape__00000049_.OverideLayoutDirection = false
	__ProductShape__00000049_.LayoutDirection = models.Vertical
	__ProductShape__00000049_.X = 86.516205
	__ProductShape__00000049_.Y = 1636.119696
	__ProductShape__00000049_.Width = 250.000000
	__ProductShape__00000049_.Height = 70.000000
	__ProductShape__00000049_.IsHidden = false

	__ProductShape__00000050_.Name = `-PIT Report`
	__ProductShape__00000050_.OverideLayoutDirection = false
	__ProductShape__00000050_.LayoutDirection = models.Vertical
	__ProductShape__00000050_.X = 86.516205
	__ProductShape__00000050_.Y = 1776.119696
	__ProductShape__00000050_.Width = 250.000000
	__ProductShape__00000050_.Height = 196.000000
	__ProductShape__00000050_.IsHidden = false

	__ProductShape__00000051_.Name = `-PIT Report`
	__ProductShape__00000051_.OverideLayoutDirection = false
	__ProductShape__00000051_.LayoutDirection = models.Vertical
	__ProductShape__00000051_.X = 99.516205
	__ProductShape__00000051_.Y = 2954.119696
	__ProductShape__00000051_.Width = 250.000000
	__ProductShape__00000051_.Height = 70.000000
	__ProductShape__00000051_.IsHidden = false

	__ProductShape__00000052_.Name = `-PIT Report`
	__ProductShape__00000052_.OverideLayoutDirection = false
	__ProductShape__00000052_.LayoutDirection = models.Vertical
	__ProductShape__00000052_.X = 99.516205
	__ProductShape__00000052_.Y = 3094.119696
	__ProductShape__00000052_.Width = 250.000000
	__ProductShape__00000052_.Height = 70.000000
	__ProductShape__00000052_.IsHidden = false

	__Resource__00000000_.Name = `Program Investigation Team (PIT)`
	__Resource__00000000_.ComputedPrefix = `1.2`
	__Resource__00000000_.IsExpanded = false
	__Resource__00000000_.IsImport = false
	__Resource__00000000_.Description = ``

	__Resource__00000001_.Name = `Barry "Butch" Wilmore`
	__Resource__00000001_.ComputedPrefix = `1.1.1.1`
	__Resource__00000001_.IsExpanded = false
	__Resource__00000001_.IsImport = false
	__Resource__00000001_.Description = ``

	__Resource__00000002_.Name = `Sunita "Suni" Williams`
	__Resource__00000002_.ComputedPrefix = `1.1.1.2`
	__Resource__00000002_.IsExpanded = false
	__Resource__00000002_.IsImport = false
	__Resource__00000002_.Description = ``

	__Resource__00000003_.Name = `NASA`
	__Resource__00000003_.ComputedPrefix = `1`
	__Resource__00000003_.IsExpanded = false
	__Resource__00000003_.IsImport = false
	__Resource__00000003_.Description = ``

	__Resource__00000004_.Name = `Crew Commercial Program (CPP2)`
	__Resource__00000004_.ComputedPrefix = `1.1`
	__Resource__00000004_.IsExpanded = false
	__Resource__00000004_.IsImport = false
	__Resource__00000004_.Description = ``

	__Resource__00000005_.Name = `Crews`
	__Resource__00000005_.ComputedPrefix = `1.1.1`
	__Resource__00000005_.IsExpanded = false
	__Resource__00000005_.IsImport = false
	__Resource__00000005_.Description = ``

	__Resource__00000006_.Name = `Boeing`
	__Resource__00000006_.ComputedPrefix = `2`
	__Resource__00000006_.IsExpanded = false
	__Resource__00000006_.IsImport = false
	__Resource__00000006_.Description = ``

	__Resource__00000007_.Name = ``
	__Resource__00000007_.ComputedPrefix = `2.1`
	__Resource__00000007_.IsExpanded = false
	__Resource__00000007_.IsImport = false
	__Resource__00000007_.Description = ``

	__Resource__00000009_.Name = ` Starliner Tests and Anomalies Review (STAR) Investigation Team`
	__Resource__00000009_.ComputedPrefix = `1.1.2`
	__Resource__00000009_.IsExpanded = false
	__Resource__00000009_.IsImport = false
	__Resource__00000009_.Description = ``

	__Resource__00000010_.Name = `Flight Control Team (FCT)`
	__Resource__00000010_.ComputedPrefix = `1.1.3`
	__Resource__00000010_.IsExpanded = false
	__Resource__00000010_.IsImport = false
	__Resource__00000010_.Description = ``

	__ResourceCompositionShape__00000000_.Name = `NASA to `
	__ResourceCompositionShape__00000000_.StartRatio = 0.500000
	__ResourceCompositionShape__00000000_.EndRatio = 0.500000
	__ResourceCompositionShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000000_.CornerOffsetRatio = 1.680000
	__ResourceCompositionShape__00000000_.IsHidden = false

	__ResourceCompositionShape__00000001_.Name = `Crew Commercial Program (CPP) to `
	__ResourceCompositionShape__00000001_.StartRatio = 0.500000
	__ResourceCompositionShape__00000001_.EndRatio = 0.500000
	__ResourceCompositionShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000001_.CornerOffsetRatio = 1.680000
	__ResourceCompositionShape__00000001_.IsHidden = false

	__ResourceCompositionShape__00000002_.Name = `Crews to Barry "Butch" Wilmore`
	__ResourceCompositionShape__00000002_.StartRatio = 0.500000
	__ResourceCompositionShape__00000002_.EndRatio = 0.500000
	__ResourceCompositionShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000002_.CornerOffsetRatio = 1.680000
	__ResourceCompositionShape__00000002_.IsHidden = false

	__ResourceCompositionShape__00000003_.Name = `Crews to Sunita "Suni" Williams`
	__ResourceCompositionShape__00000003_.StartRatio = 0.500000
	__ResourceCompositionShape__00000003_.EndRatio = 0.500000
	__ResourceCompositionShape__00000003_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000003_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000003_.CornerOffsetRatio = 1.680000
	__ResourceCompositionShape__00000003_.IsHidden = false

	__ResourceCompositionShape__00000004_.Name = `NASA to PITProgram Investigation Team (PIT)`
	__ResourceCompositionShape__00000004_.StartRatio = 0.500000
	__ResourceCompositionShape__00000004_.EndRatio = 0.500000
	__ResourceCompositionShape__00000004_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000004_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000004_.CornerOffsetRatio = 1.680000
	__ResourceCompositionShape__00000004_.IsHidden = false

	__ResourceCompositionShape__00000005_.Name = `Boeing to `
	__ResourceCompositionShape__00000005_.StartRatio = 0.500000
	__ResourceCompositionShape__00000005_.EndRatio = 0.500000
	__ResourceCompositionShape__00000005_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000005_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000005_.CornerOffsetRatio = 1.680000
	__ResourceCompositionShape__00000005_.IsHidden = false

	__ResourceCompositionShape__00000007_.Name = `Crew Commercial Program (CPP) to `
	__ResourceCompositionShape__00000007_.StartRatio = 0.500000
	__ResourceCompositionShape__00000007_.EndRatio = 0.500000
	__ResourceCompositionShape__00000007_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000007_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000007_.CornerOffsetRatio = 1.680000
	__ResourceCompositionShape__00000007_.IsHidden = false

	__ResourceCompositionShape__00000008_.Name = `Crew Commercial Program (CPP) to `
	__ResourceCompositionShape__00000008_.StartRatio = 0.500000
	__ResourceCompositionShape__00000008_.EndRatio = 0.500000
	__ResourceCompositionShape__00000008_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000008_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000008_.CornerOffsetRatio = 1.680000
	__ResourceCompositionShape__00000008_.IsHidden = false

	__ResourceShape__00000000_.Name = `-Default Diagram`
	__ResourceShape__00000000_.X = 52.114853
	__ResourceShape__00000000_.Y = 34.127119
	__ResourceShape__00000000_.Width = 250.000000
	__ResourceShape__00000000_.Height = 70.000000
	__ResourceShape__00000000_.IsHidden = false

	__ResourceShape__00000001_.Name = `-Default Diagram`
	__ResourceShape__00000001_.X = 34.180389
	__ResourceShape__00000001_.Y = 553.010316
	__ResourceShape__00000001_.Width = 250.000000
	__ResourceShape__00000001_.Height = 70.000000
	__ResourceShape__00000001_.IsHidden = false

	__ResourceShape__00000002_.Name = `-Default Diagram`
	__ResourceShape__00000002_.X = 33.201744
	__ResourceShape__00000002_.Y = 650.719328
	__ResourceShape__00000002_.Width = 250.000000
	__ResourceShape__00000002_.Height = 70.000000
	__ResourceShape__00000002_.IsHidden = false

	__ResourceShape__00000003_.Name = `-RBS`
	__ResourceShape__00000003_.X = 134.737820
	__ResourceShape__00000003_.Y = 123.124371
	__ResourceShape__00000003_.Width = 250.000000
	__ResourceShape__00000003_.Height = 70.000000
	__ResourceShape__00000003_.IsHidden = false

	__ResourceShape__00000004_.Name = `-RBS`
	__ResourceShape__00000004_.X = 137.737820
	__ResourceShape__00000004_.Y = 300.124371
	__ResourceShape__00000004_.Width = 250.000000
	__ResourceShape__00000004_.Height = 70.000000
	__ResourceShape__00000004_.IsHidden = false

	__ResourceShape__00000005_.Name = `-RBS`
	__ResourceShape__00000005_.X = 137.737820
	__ResourceShape__00000005_.Y = 440.124371
	__ResourceShape__00000005_.Width = 250.000000
	__ResourceShape__00000005_.Height = 70.000000
	__ResourceShape__00000005_.IsHidden = false

	__ResourceShape__00000006_.Name = `Barry "Butch" Wilmore-RBS`
	__ResourceShape__00000006_.X = 132.423066
	__ResourceShape__00000006_.Y = 597.719206
	__ResourceShape__00000006_.Width = 250.000000
	__ResourceShape__00000006_.Height = 70.000000
	__ResourceShape__00000006_.IsHidden = false

	__ResourceShape__00000007_.Name = `Sunita "Suni" Williams-RBS`
	__ResourceShape__00000007_.X = 454.596590
	__ResourceShape__00000007_.Y = 600.173621
	__ResourceShape__00000007_.Width = 250.000000
	__ResourceShape__00000007_.Height = 70.000000
	__ResourceShape__00000007_.IsHidden = false

	__ResourceShape__00000008_.Name = `PITProgram Investigation Team (PIT)-RBS`
	__ResourceShape__00000008_.X = 835.651635
	__ResourceShape__00000008_.Y = 297.886860
	__ResourceShape__00000008_.Width = 250.000000
	__ResourceShape__00000008_.Height = 70.000000
	__ResourceShape__00000008_.IsHidden = false

	__ResourceShape__00000009_.Name = `-RBS`
	__ResourceShape__00000009_.X = 126.434105
	__ResourceShape__00000009_.Y = 846.201664
	__ResourceShape__00000009_.Width = 250.000000
	__ResourceShape__00000009_.Height = 70.000000
	__ResourceShape__00000009_.IsHidden = false

	__ResourceShape__00000010_.Name = `-RBS`
	__ResourceShape__00000010_.X = 110.434105
	__ResourceShape__00000010_.Y = 1029.201740
	__ResourceShape__00000010_.Width = 250.000000
	__ResourceShape__00000010_.Height = 70.000000
	__ResourceShape__00000010_.IsHidden = false

	__ResourceShape__00000012_.Name = `-RBS`
	__ResourceShape__00000012_.X = 437.737820
	__ResourceShape__00000012_.Y = 440.124371
	__ResourceShape__00000012_.Width = 413.000000
	__ResourceShape__00000012_.Height = 70.000000
	__ResourceShape__00000012_.IsHidden = false

	__ResourceShape__00000013_.Name = `Program Investigation Team (PIT)-PBS`
	__ResourceShape__00000013_.X = 155.000061
	__ResourceShape__00000013_.Y = 835.999985
	__ResourceShape__00000013_.Width = 250.000000
	__ResourceShape__00000013_.Height = 70.000000
	__ResourceShape__00000013_.IsHidden = false

	__ResourceShape__00000014_.Name = `Program Investigation Team (PIT)-PIT focus`
	__ResourceShape__00000014_.X = 104.267926
	__ResourceShape__00000014_.Y = 288.658205
	__ResourceShape__00000014_.Width = 250.000000
	__ResourceShape__00000014_.Height = 70.000000
	__ResourceShape__00000014_.IsHidden = false

	__ResourceShape__00000015_.Name = `-RBS`
	__ResourceShape__00000015_.X = 883.737820
	__ResourceShape__00000015_.Y = 438.124371
	__ResourceShape__00000015_.Width = 250.000000
	__ResourceShape__00000015_.Height = 70.000000
	__ResourceShape__00000015_.IsHidden = false

	__ResourceTaskShape__00000000_.Name = `PITProgram Investigation Team (PIT) to Mishap investigation`
	__ResourceTaskShape__00000000_.StartRatio = 0.500000
	__ResourceTaskShape__00000000_.EndRatio = 0.500000
	__ResourceTaskShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceTaskShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceTaskShape__00000000_.CornerOffsetRatio = 1.680000
	__ResourceTaskShape__00000000_.IsHidden = false

	__Task__00000000_.Name = `Starliner Crewed Flight Test (CFT)`
	__Task__00000000_.ComputedPrefix = `2.3`
	__Task__00000000_.IsExpanded = false
	__Task__00000000_.IsImport = false
	__Task__00000000_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2024-06-05 00:00:00 +0000 UTC")
	__Task__00000000_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2024-06-10 00:00:00 +0000 UTC")
	__Task__00000000_.Description = `The mission of interest to the report.`
	__Task__00000000_.IsInputsNodeExpanded = false
	__Task__00000000_.IsOutputsNodeExpanded = false
	__Task__00000000_.IsWithCompletion = false
	__Task__00000000_.Completion = ""

	__Task__00000001_.Name = `Mishap investigations`
	__Task__00000001_.ComputedPrefix = `1`
	__Task__00000001_.IsExpanded = false
	__Task__00000001_.IsImport = false
	__Task__00000001_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2025-02-01 00:00:00 +0000 UTC")
	__Task__00000001_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2025-11-01 00:00:00 +0000 UTC")
	__Task__00000001_.Description = ``
	__Task__00000001_.IsInputsNodeExpanded = false
	__Task__00000001_.IsOutputsNodeExpanded = false
	__Task__00000001_.IsWithCompletion = false
	__Task__00000001_.Completion = ""

	__Task__00000002_.Name = `Commercial Crew Program (CCP),`
	__Task__00000002_.ComputedPrefix = `2`
	__Task__00000002_.IsExpanded = false
	__Task__00000002_.IsImport = false
	__Task__00000002_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000002_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000002_.Description = ``
	__Task__00000002_.IsInputsNodeExpanded = false
	__Task__00000002_.IsOutputsNodeExpanded = false
	__Task__00000002_.IsWithCompletion = false
	__Task__00000002_.Completion = ""

	__Task__00000003_.Name = ` Commercial ReSupply (CRS) `
	__Task__00000003_.ComputedPrefix = `3`
	__Task__00000003_.IsExpanded = false
	__Task__00000003_.IsImport = false
	__Task__00000003_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000003_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000003_.Description = ``
	__Task__00000003_.IsInputsNodeExpanded = false
	__Task__00000003_.IsOutputsNodeExpanded = false
	__Task__00000003_.IsWithCompletion = false
	__Task__00000003_.Completion = ""

	__Task__00000004_.Name = `Program Investigation Team (PIT) Report`
	__Task__00000004_.ComputedPrefix = `1.1`
	__Task__00000004_.IsExpanded = false
	__Task__00000004_.IsImport = false
	__Task__00000004_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000004_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000004_.Description = ``
	__Task__00000004_.IsInputsNodeExpanded = false
	__Task__00000004_.IsOutputsNodeExpanded = false
	__Task__00000004_.IsWithCompletion = false
	__Task__00000004_.Completion = ""

	__Task__00000005_.Name = `Orbital Flight Test-1 (OFT-1)`
	__Task__00000005_.ComputedPrefix = `2.1`
	__Task__00000005_.IsExpanded = false
	__Task__00000005_.IsImport = false
	__Task__00000005_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2019-12-20 00:00:00 +0000 UTC")
	__Task__00000005_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2019-12-20 00:00:00 +0000 UTC")
	__Task__00000005_.Description = ``
	__Task__00000005_.IsInputsNodeExpanded = false
	__Task__00000005_.IsOutputsNodeExpanded = false
	__Task__00000005_.IsWithCompletion = false
	__Task__00000005_.Completion = ""

	__Task__00000006_.Name = `Orbital Flight Test-2 (OFT-2)`
	__Task__00000006_.ComputedPrefix = `2.2`
	__Task__00000006_.IsExpanded = false
	__Task__00000006_.IsImport = false
	__Task__00000006_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2022-05-19 00:00:00 +0000 UTC")
	__Task__00000006_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2022-05-25 00:00:00 +0000 UTC")
	__Task__00000006_.Description = ``
	__Task__00000006_.IsInputsNodeExpanded = false
	__Task__00000006_.IsOutputsNodeExpanded = false
	__Task__00000006_.IsWithCompletion = false
	__Task__00000006_.Completion = ""

	__Task__00000007_.Name = `ISS Approach`
	__Task__00000007_.ComputedPrefix = `2.3.1`
	__Task__00000007_.IsExpanded = false
	__Task__00000007_.IsImport = false
	__Task__00000007_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2024-06-05 00:00:00 +0000 UTC")
	__Task__00000007_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2024-06-05 00:00:00 +0000 UTC")
	__Task__00000007_.Description = ``
	__Task__00000007_.IsInputsNodeExpanded = false
	__Task__00000007_.IsOutputsNodeExpanded = false
	__Task__00000007_.IsWithCompletion = false
	__Task__00000007_.Completion = ""

	__TaskCompositionShape__00000001_.Name = `Commercial Crew Program (CCP), to Starliner Crewed Flight Test (CFT)`
	__TaskCompositionShape__00000001_.StartRatio = 0.500000
	__TaskCompositionShape__00000001_.EndRatio = 0.500000
	__TaskCompositionShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000001_.CornerOffsetRatio = 1.680000
	__TaskCompositionShape__00000001_.IsHidden = false

	__TaskCompositionShape__00000002_.Name = `Mishap investigations to `
	__TaskCompositionShape__00000002_.StartRatio = 0.500000
	__TaskCompositionShape__00000002_.EndRatio = 0.500000
	__TaskCompositionShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000002_.CornerOffsetRatio = 1.680000
	__TaskCompositionShape__00000002_.IsHidden = false

	__TaskCompositionShape__00000003_.Name = `Commercial Crew Program (CCP), to `
	__TaskCompositionShape__00000003_.StartRatio = 0.500000
	__TaskCompositionShape__00000003_.EndRatio = 0.500000
	__TaskCompositionShape__00000003_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000003_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000003_.CornerOffsetRatio = 1.680000
	__TaskCompositionShape__00000003_.IsHidden = false

	__TaskCompositionShape__00000004_.Name = `Commercial Crew Program (CCP), to `
	__TaskCompositionShape__00000004_.StartRatio = 0.500000
	__TaskCompositionShape__00000004_.EndRatio = 0.500000
	__TaskCompositionShape__00000004_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000004_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000004_.CornerOffsetRatio = 1.680000
	__TaskCompositionShape__00000004_.IsHidden = false

	__TaskCompositionShape__00000005_.Name = `Starliner Crewed Flight Test (CFT) to `
	__TaskCompositionShape__00000005_.StartRatio = 0.500000
	__TaskCompositionShape__00000005_.EndRatio = 0.500000
	__TaskCompositionShape__00000005_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000005_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000005_.CornerOffsetRatio = 1.680000
	__TaskCompositionShape__00000005_.IsHidden = false

	__TaskInputShape__00000000_.Name = `Program Investigation Team (PIT) Report to  Starliner Tests and Anomalies Review (STAR) Investigation Report`
	__TaskInputShape__00000000_.StartRatio = 0.278248
	__TaskInputShape__00000000_.EndRatio = 0.277627
	__TaskInputShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskInputShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskInputShape__00000000_.CornerOffsetRatio = 1.883372
	__TaskInputShape__00000000_.IsHidden = false

	__TaskOutputShape__00000000_.Name = `Program Investigation Team (PIT) Report to Program Investigation Team (PIT) Report`
	__TaskOutputShape__00000000_.StartRatio = 0.439622
	__TaskOutputShape__00000000_.EndRatio = 0.501995
	__TaskOutputShape__00000000_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000000_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000000_.CornerOffsetRatio = 1.313627
	__TaskOutputShape__00000000_.IsHidden = false

	__TaskShape__00000003_.Name = `Mishap investigation-WBS`
	__TaskShape__00000003_.IsShowDate = false
	__TaskShape__00000003_.X = 28.416829
	__TaskShape__00000003_.Y = 546.674065
	__TaskShape__00000003_.Width = 250.000000
	__TaskShape__00000003_.Height = 70.000000
	__TaskShape__00000003_.IsHidden = false

	__TaskShape__00000004_.Name = `Commercial Crew Program (CCP),-WBS`
	__TaskShape__00000004_.IsShowDate = false
	__TaskShape__00000004_.X = 14.071983
	__TaskShape__00000004_.Y = 166.517160
	__TaskShape__00000004_.Width = 250.000000
	__TaskShape__00000004_.Height = 70.000000
	__TaskShape__00000004_.IsHidden = false

	__TaskShape__00000005_.Name = `Starliner Crewed Flight Test (CFT)-WBS`
	__TaskShape__00000005_.IsShowDate = false
	__TaskShape__00000005_.X = 645.617766
	__TaskShape__00000005_.Y = 360.295606
	__TaskShape__00000005_.Width = 250.000000
	__TaskShape__00000005_.Height = 70.000000
	__TaskShape__00000005_.IsHidden = false

	__TaskShape__00000006_.Name = `-WBS`
	__TaskShape__00000006_.IsShowDate = false
	__TaskShape__00000006_.X = 14.409715
	__TaskShape__00000006_.Y = 32.470060
	__TaskShape__00000006_.Width = 250.000000
	__TaskShape__00000006_.Height = 70.000000
	__TaskShape__00000006_.IsHidden = false

	__TaskShape__00000007_.Name = `Mishap investigation-PBS`
	__TaskShape__00000007_.IsShowDate = false
	__TaskShape__00000007_.X = 117.000061
	__TaskShape__00000007_.Y = 743.999985
	__TaskShape__00000007_.Width = 250.000000
	__TaskShape__00000007_.Height = 70.000000
	__TaskShape__00000007_.IsHidden = false

	__TaskShape__00000008_.Name = `-WBS`
	__TaskShape__00000008_.IsShowDate = false
	__TaskShape__00000008_.X = 28.416829
	__TaskShape__00000008_.Y = 686.674065
	__TaskShape__00000008_.Width = 250.000000
	__TaskShape__00000008_.Height = 70.000000
	__TaskShape__00000008_.IsHidden = false

	__TaskShape__00000009_.Name = `Program Investigation Team (PIT) Report-PIT focus`
	__TaskShape__00000009_.IsShowDate = false
	__TaskShape__00000009_.X = 499.174738
	__TaskShape__00000009_.Y = 283.233412
	__TaskShape__00000009_.Width = 250.000000
	__TaskShape__00000009_.Height = 70.000000
	__TaskShape__00000009_.IsHidden = false

	__TaskShape__00000010_.Name = `-WBS`
	__TaskShape__00000010_.IsShowDate = false
	__TaskShape__00000010_.X = 19.071983
	__TaskShape__00000010_.Y = 346.517160
	__TaskShape__00000010_.Width = 250.000000
	__TaskShape__00000010_.Height = 70.000000
	__TaskShape__00000010_.IsHidden = false

	__TaskShape__00000011_.Name = `-WBS`
	__TaskShape__00000011_.IsShowDate = false
	__TaskShape__00000011_.X = 336.071983
	__TaskShape__00000011_.Y = 355.517160
	__TaskShape__00000011_.Width = 250.000000
	__TaskShape__00000011_.Height = 70.000000
	__TaskShape__00000011_.IsHidden = false

	__TaskShape__00000012_.Name = `-WBS`
	__TaskShape__00000012_.IsShowDate = false
	__TaskShape__00000012_.X = 645.617766
	__TaskShape__00000012_.Y = 500.295606
	__TaskShape__00000012_.Width = 250.000000
	__TaskShape__00000012_.Height = 70.000000
	__TaskShape__00000012_.IsHidden = false

	// insertion point for setup of pointers
	__Diagram__00000001_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000001_.ProductsWhoseNodeIsExpanded, __Product__00000007_)
	__Diagram__00000001_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000001_.ProductsWhoseNodeIsExpanded, __Product__00000006_)
	__Diagram__00000001_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000001_.ProductsWhoseNodeIsExpanded, __Product__00000002_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000003_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000004_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000005_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000006_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000008_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000010_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000011_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000012_)
	__Diagram__00000001_.TasksWhoseNodeIsExpanded = append(__Diagram__00000001_.TasksWhoseNodeIsExpanded, __Task__00000002_)
	__Diagram__00000001_.TasksWhoseNodeIsExpanded = append(__Diagram__00000001_.TasksWhoseNodeIsExpanded, __Task__00000001_)
	__Diagram__00000001_.TasksWhoseNodeIsExpanded = append(__Diagram__00000001_.TasksWhoseNodeIsExpanded, __Task__00000000_)
	__Diagram__00000001_.TaskComposition_Shapes = append(__Diagram__00000001_.TaskComposition_Shapes, __TaskCompositionShape__00000001_)
	__Diagram__00000001_.TaskComposition_Shapes = append(__Diagram__00000001_.TaskComposition_Shapes, __TaskCompositionShape__00000002_)
	__Diagram__00000001_.TaskComposition_Shapes = append(__Diagram__00000001_.TaskComposition_Shapes, __TaskCompositionShape__00000003_)
	__Diagram__00000001_.TaskComposition_Shapes = append(__Diagram__00000001_.TaskComposition_Shapes, __TaskCompositionShape__00000004_)
	__Diagram__00000001_.TaskComposition_Shapes = append(__Diagram__00000001_.TaskComposition_Shapes, __TaskCompositionShape__00000005_)
	__Diagram__00000001_.Note_Shapes = append(__Diagram__00000001_.Note_Shapes, __NoteShape__00000017_)
	__Diagram__00000001_.NoteTaskShapes = append(__Diagram__00000001_.NoteTaskShapes, __NoteTaskShape__00000001_)
	__Diagram__00000001_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000001_.ResourcesWhoseNodeIsExpanded, __Resource__00000003_)
	__Diagram__00000001_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000001_.ResourcesWhoseNodeIsExpanded, __Resource__00000004_)
	__Diagram__00000001_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000001_.ResourcesWhoseNodeIsExpanded, __Resource__00000000_)
	__Diagram__00000002_.Product_Shapes = append(__Diagram__00000002_.Product_Shapes, __ProductShape__00000002_)
	__Diagram__00000002_.Product_Shapes = append(__Diagram__00000002_.Product_Shapes, __ProductShape__00000003_)
	__Diagram__00000002_.Product_Shapes = append(__Diagram__00000002_.Product_Shapes, __ProductShape__00000004_)
	__Diagram__00000002_.Product_Shapes = append(__Diagram__00000002_.Product_Shapes, __ProductShape__00000005_)
	__Diagram__00000002_.Product_Shapes = append(__Diagram__00000002_.Product_Shapes, __ProductShape__00000006_)
	__Diagram__00000002_.Product_Shapes = append(__Diagram__00000002_.Product_Shapes, __ProductShape__00000009_)
	__Diagram__00000002_.Product_Shapes = append(__Diagram__00000002_.Product_Shapes, __ProductShape__00000010_)
	__Diagram__00000002_.Product_Shapes = append(__Diagram__00000002_.Product_Shapes, __ProductShape__00000011_)
	__Diagram__00000002_.Product_Shapes = append(__Diagram__00000002_.Product_Shapes, __ProductShape__00000027_)
	__Diagram__00000002_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000002_.ProductsWhoseNodeIsExpanded, __Product__00000006_)
	__Diagram__00000002_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000002_.ProductsWhoseNodeIsExpanded, __Product__00000002_)
	__Diagram__00000002_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000002_.ProductsWhoseNodeIsExpanded, __Product__00000005_)
	__Diagram__00000002_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000002_.ProductsWhoseNodeIsExpanded, __Product__00000017_)
	__Diagram__00000002_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000002_.ProductsWhoseNodeIsExpanded, __Product__00000003_)
	__Diagram__00000002_.ProductComposition_Shapes = append(__Diagram__00000002_.ProductComposition_Shapes, __ProductCompositionShape__00000000_)
	__Diagram__00000002_.ProductComposition_Shapes = append(__Diagram__00000002_.ProductComposition_Shapes, __ProductCompositionShape__00000002_)
	__Diagram__00000002_.ProductComposition_Shapes = append(__Diagram__00000002_.ProductComposition_Shapes, __ProductCompositionShape__00000003_)
	__Diagram__00000002_.ProductComposition_Shapes = append(__Diagram__00000002_.ProductComposition_Shapes, __ProductCompositionShape__00000004_)
	__Diagram__00000002_.ProductComposition_Shapes = append(__Diagram__00000002_.ProductComposition_Shapes, __ProductCompositionShape__00000005_)
	__Diagram__00000002_.ProductComposition_Shapes = append(__Diagram__00000002_.ProductComposition_Shapes, __ProductCompositionShape__00000018_)
	__Diagram__00000002_.ProductComposition_Shapes = append(__Diagram__00000002_.ProductComposition_Shapes, __ProductCompositionShape__00000044_)
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
	__Diagram__00000003_.Resource_Shapes = append(__Diagram__00000003_.Resource_Shapes, __ResourceShape__00000012_)
	__Diagram__00000003_.Resource_Shapes = append(__Diagram__00000003_.Resource_Shapes, __ResourceShape__00000015_)
	__Diagram__00000003_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000003_.ResourcesWhoseNodeIsExpanded, __Resource__00000000_)
	__Diagram__00000003_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000003_.ResourcesWhoseNodeIsExpanded, __Resource__00000006_)
	__Diagram__00000003_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000003_.ResourcesWhoseNodeIsExpanded, __Resource__00000005_)
	__Diagram__00000003_.ResourceComposition_Shapes = append(__Diagram__00000003_.ResourceComposition_Shapes, __ResourceCompositionShape__00000000_)
	__Diagram__00000003_.ResourceComposition_Shapes = append(__Diagram__00000003_.ResourceComposition_Shapes, __ResourceCompositionShape__00000001_)
	__Diagram__00000003_.ResourceComposition_Shapes = append(__Diagram__00000003_.ResourceComposition_Shapes, __ResourceCompositionShape__00000002_)
	__Diagram__00000003_.ResourceComposition_Shapes = append(__Diagram__00000003_.ResourceComposition_Shapes, __ResourceCompositionShape__00000003_)
	__Diagram__00000003_.ResourceComposition_Shapes = append(__Diagram__00000003_.ResourceComposition_Shapes, __ResourceCompositionShape__00000004_)
	__Diagram__00000003_.ResourceComposition_Shapes = append(__Diagram__00000003_.ResourceComposition_Shapes, __ResourceCompositionShape__00000005_)
	__Diagram__00000003_.ResourceComposition_Shapes = append(__Diagram__00000003_.ResourceComposition_Shapes, __ResourceCompositionShape__00000007_)
	__Diagram__00000003_.ResourceComposition_Shapes = append(__Diagram__00000003_.ResourceComposition_Shapes, __ResourceCompositionShape__00000008_)
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
	__Diagram__00000004_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000004_.ResourcesWhoseNodeIsExpanded, __Resource__00000005_)
	__Diagram__00000004_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000004_.ResourcesWhoseNodeIsExpanded, __Resource__00000004_)
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000012_)
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000013_)
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000014_)
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000022_)
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000023_)
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000028_)
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000029_)
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000032_)
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000034_)
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000035_)
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000036_)
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000037_)
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000039_)
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000040_)
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000041_)
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000042_)
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000045_)
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000046_)
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000047_)
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000048_)
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000049_)
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000050_)
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000051_)
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000052_)
	__Diagram__00000005_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000005_.ProductsWhoseNodeIsExpanded, __Product__00000006_)
	__Diagram__00000005_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000005_.ProductsWhoseNodeIsExpanded, __Product__00000017_)
	__Diagram__00000005_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000005_.ProductsWhoseNodeIsExpanded, __Product__00000009_)
	__Diagram__00000005_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000005_.ProductsWhoseNodeIsExpanded, __Product__00000029_)
	__Diagram__00000005_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000005_.ProductsWhoseNodeIsExpanded, __Product__00000005_)
	__Diagram__00000005_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000005_.ProductsWhoseNodeIsExpanded, __Product__00000035_)
	__Diagram__00000005_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000005_.ProductsWhoseNodeIsExpanded, __Product__00000045_)
	__Diagram__00000005_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000005_.ProductsWhoseNodeIsExpanded, __Product__00000003_)
	__Diagram__00000005_.ProductComposition_Shapes = append(__Diagram__00000005_.ProductComposition_Shapes, __ProductCompositionShape__00000006_)
	__Diagram__00000005_.ProductComposition_Shapes = append(__Diagram__00000005_.ProductComposition_Shapes, __ProductCompositionShape__00000013_)
	__Diagram__00000005_.ProductComposition_Shapes = append(__Diagram__00000005_.ProductComposition_Shapes, __ProductCompositionShape__00000014_)
	__Diagram__00000005_.ProductComposition_Shapes = append(__Diagram__00000005_.ProductComposition_Shapes, __ProductCompositionShape__00000019_)
	__Diagram__00000005_.ProductComposition_Shapes = append(__Diagram__00000005_.ProductComposition_Shapes, __ProductCompositionShape__00000020_)
	__Diagram__00000005_.ProductComposition_Shapes = append(__Diagram__00000005_.ProductComposition_Shapes, __ProductCompositionShape__00000023_)
	__Diagram__00000005_.ProductComposition_Shapes = append(__Diagram__00000005_.ProductComposition_Shapes, __ProductCompositionShape__00000025_)
	__Diagram__00000005_.ProductComposition_Shapes = append(__Diagram__00000005_.ProductComposition_Shapes, __ProductCompositionShape__00000026_)
	__Diagram__00000005_.ProductComposition_Shapes = append(__Diagram__00000005_.ProductComposition_Shapes, __ProductCompositionShape__00000027_)
	__Diagram__00000005_.ProductComposition_Shapes = append(__Diagram__00000005_.ProductComposition_Shapes, __ProductCompositionShape__00000028_)
	__Diagram__00000005_.ProductComposition_Shapes = append(__Diagram__00000005_.ProductComposition_Shapes, __ProductCompositionShape__00000030_)
	__Diagram__00000005_.ProductComposition_Shapes = append(__Diagram__00000005_.ProductComposition_Shapes, __ProductCompositionShape__00000031_)
	__Diagram__00000005_.ProductComposition_Shapes = append(__Diagram__00000005_.ProductComposition_Shapes, __ProductCompositionShape__00000032_)
	__Diagram__00000005_.ProductComposition_Shapes = append(__Diagram__00000005_.ProductComposition_Shapes, __ProductCompositionShape__00000033_)
	__Diagram__00000005_.ProductComposition_Shapes = append(__Diagram__00000005_.ProductComposition_Shapes, __ProductCompositionShape__00000036_)
	__Diagram__00000005_.ProductComposition_Shapes = append(__Diagram__00000005_.ProductComposition_Shapes, __ProductCompositionShape__00000037_)
	__Diagram__00000005_.ProductComposition_Shapes = append(__Diagram__00000005_.ProductComposition_Shapes, __ProductCompositionShape__00000038_)
	__Diagram__00000005_.ProductComposition_Shapes = append(__Diagram__00000005_.ProductComposition_Shapes, __ProductCompositionShape__00000039_)
	__Diagram__00000005_.ProductComposition_Shapes = append(__Diagram__00000005_.ProductComposition_Shapes, __ProductCompositionShape__00000040_)
	__Diagram__00000005_.ProductComposition_Shapes = append(__Diagram__00000005_.ProductComposition_Shapes, __ProductCompositionShape__00000041_)
	__Diagram__00000005_.ProductComposition_Shapes = append(__Diagram__00000005_.ProductComposition_Shapes, __ProductCompositionShape__00000042_)
	__Diagram__00000005_.ProductComposition_Shapes = append(__Diagram__00000005_.ProductComposition_Shapes, __ProductCompositionShape__00000043_)
	__Diagram__00000005_.Note_Shapes = append(__Diagram__00000005_.Note_Shapes, __NoteShape__00000003_)
	__Diagram__00000005_.Note_Shapes = append(__Diagram__00000005_.Note_Shapes, __NoteShape__00000004_)
	__Diagram__00000005_.Note_Shapes = append(__Diagram__00000005_.Note_Shapes, __NoteShape__00000013_)
	__Diagram__00000005_.Note_Shapes = append(__Diagram__00000005_.Note_Shapes, __NoteShape__00000014_)
	__Diagram__00000005_.Note_Shapes = append(__Diagram__00000005_.Note_Shapes, __NoteShape__00000015_)
	__Diagram__00000005_.Note_Shapes = append(__Diagram__00000005_.Note_Shapes, __NoteShape__00000016_)
	__Diagram__00000005_.Note_Shapes = append(__Diagram__00000005_.Note_Shapes, __NoteShape__00000018_)
	__Diagram__00000005_.Note_Shapes = append(__Diagram__00000005_.Note_Shapes, __NoteShape__00000019_)
	__Diagram__00000005_.Note_Shapes = append(__Diagram__00000005_.Note_Shapes, __NoteShape__00000020_)
	__Diagram__00000005_.Note_Shapes = append(__Diagram__00000005_.Note_Shapes, __NoteShape__00000021_)
	__Diagram__00000005_.Note_Shapes = append(__Diagram__00000005_.Note_Shapes, __NoteShape__00000022_)
	__Diagram__00000005_.Note_Shapes = append(__Diagram__00000005_.Note_Shapes, __NoteShape__00000023_)
	__Diagram__00000005_.Note_Shapes = append(__Diagram__00000005_.Note_Shapes, __NoteShape__00000024_)
	__Diagram__00000005_.Note_Shapes = append(__Diagram__00000005_.Note_Shapes, __NoteShape__00000025_)
	__Diagram__00000005_.NoteProductShapes = append(__Diagram__00000005_.NoteProductShapes, __NoteProductShape__00000002_)
	__Diagram__00000005_.NoteProductShapes = append(__Diagram__00000005_.NoteProductShapes, __NoteProductShape__00000003_)
	__Diagram__00000005_.NoteProductShapes = append(__Diagram__00000005_.NoteProductShapes, __NoteProductShape__00000016_)
	__Diagram__00000005_.NoteProductShapes = append(__Diagram__00000005_.NoteProductShapes, __NoteProductShape__00000017_)
	__Diagram__00000005_.NoteProductShapes = append(__Diagram__00000005_.NoteProductShapes, __NoteProductShape__00000018_)
	__Diagram__00000005_.NoteProductShapes = append(__Diagram__00000005_.NoteProductShapes, __NoteProductShape__00000019_)
	__Diagram__00000005_.NoteProductShapes = append(__Diagram__00000005_.NoteProductShapes, __NoteProductShape__00000020_)
	__Diagram__00000005_.NoteProductShapes = append(__Diagram__00000005_.NoteProductShapes, __NoteProductShape__00000021_)
	__Diagram__00000005_.NoteProductShapes = append(__Diagram__00000005_.NoteProductShapes, __NoteProductShape__00000022_)
	__Diagram__00000005_.NoteProductShapes = append(__Diagram__00000005_.NoteProductShapes, __NoteProductShape__00000023_)
	__Diagram__00000005_.NoteProductShapes = append(__Diagram__00000005_.NoteProductShapes, __NoteProductShape__00000024_)
	__Diagram__00000005_.NoteProductShapes = append(__Diagram__00000005_.NoteProductShapes, __NoteProductShape__00000025_)
	__Diagram__00000005_.NoteProductShapes = append(__Diagram__00000005_.NoteProductShapes, __NoteProductShape__00000026_)
	__Diagram__00000006_.Product_Shapes = append(__Diagram__00000006_.Product_Shapes, __ProductShape__00000015_)
	__Diagram__00000006_.Product_Shapes = append(__Diagram__00000006_.Product_Shapes, __ProductShape__00000016_)
	__Diagram__00000006_.Product_Shapes = append(__Diagram__00000006_.Product_Shapes, __ProductShape__00000017_)
	__Diagram__00000006_.Product_Shapes = append(__Diagram__00000006_.Product_Shapes, __ProductShape__00000018_)
	__Diagram__00000006_.Product_Shapes = append(__Diagram__00000006_.Product_Shapes, __ProductShape__00000019_)
	__Diagram__00000006_.Product_Shapes = append(__Diagram__00000006_.Product_Shapes, __ProductShape__00000020_)
	__Diagram__00000006_.Product_Shapes = append(__Diagram__00000006_.Product_Shapes, __ProductShape__00000021_)
	__Diagram__00000006_.Product_Shapes = append(__Diagram__00000006_.Product_Shapes, __ProductShape__00000024_)
	__Diagram__00000006_.Product_Shapes = append(__Diagram__00000006_.Product_Shapes, __ProductShape__00000025_)
	__Diagram__00000006_.Product_Shapes = append(__Diagram__00000006_.Product_Shapes, __ProductShape__00000030_)
	__Diagram__00000006_.Product_Shapes = append(__Diagram__00000006_.Product_Shapes, __ProductShape__00000031_)
	__Diagram__00000006_.Product_Shapes = append(__Diagram__00000006_.Product_Shapes, __ProductShape__00000033_)
	__Diagram__00000006_.Product_Shapes = append(__Diagram__00000006_.Product_Shapes, __ProductShape__00000038_)
	__Diagram__00000006_.Product_Shapes = append(__Diagram__00000006_.Product_Shapes, __ProductShape__00000043_)
	__Diagram__00000006_.Product_Shapes = append(__Diagram__00000006_.Product_Shapes, __ProductShape__00000044_)
	__Diagram__00000006_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000006_.ProductsWhoseNodeIsExpanded, __Product__00000006_)
	__Diagram__00000006_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000006_.ProductsWhoseNodeIsExpanded, __Product__00000012_)
	__Diagram__00000006_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000006_.ProductsWhoseNodeIsExpanded, __Product__00000015_)
	__Diagram__00000006_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000006_.ProductsWhoseNodeIsExpanded, __Product__00000013_)
	__Diagram__00000006_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000006_.ProductsWhoseNodeIsExpanded, __Product__00000003_)
	__Diagram__00000006_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000006_.ProductsWhoseNodeIsExpanded, __Product__00000033_)
	__Diagram__00000006_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000006_.ProductsWhoseNodeIsExpanded, __Product__00000007_)
	__Diagram__00000006_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000006_.ProductsWhoseNodeIsExpanded, __Product__00000021_)
	__Diagram__00000006_.ProductComposition_Shapes = append(__Diagram__00000006_.ProductComposition_Shapes, __ProductCompositionShape__00000007_)
	__Diagram__00000006_.ProductComposition_Shapes = append(__Diagram__00000006_.ProductComposition_Shapes, __ProductCompositionShape__00000008_)
	__Diagram__00000006_.ProductComposition_Shapes = append(__Diagram__00000006_.ProductComposition_Shapes, __ProductCompositionShape__00000009_)
	__Diagram__00000006_.ProductComposition_Shapes = append(__Diagram__00000006_.ProductComposition_Shapes, __ProductCompositionShape__00000010_)
	__Diagram__00000006_.ProductComposition_Shapes = append(__Diagram__00000006_.ProductComposition_Shapes, __ProductCompositionShape__00000011_)
	__Diagram__00000006_.ProductComposition_Shapes = append(__Diagram__00000006_.ProductComposition_Shapes, __ProductCompositionShape__00000012_)
	__Diagram__00000006_.ProductComposition_Shapes = append(__Diagram__00000006_.ProductComposition_Shapes, __ProductCompositionShape__00000016_)
	__Diagram__00000006_.ProductComposition_Shapes = append(__Diagram__00000006_.ProductComposition_Shapes, __ProductCompositionShape__00000021_)
	__Diagram__00000006_.ProductComposition_Shapes = append(__Diagram__00000006_.ProductComposition_Shapes, __ProductCompositionShape__00000022_)
	__Diagram__00000006_.ProductComposition_Shapes = append(__Diagram__00000006_.ProductComposition_Shapes, __ProductCompositionShape__00000024_)
	__Diagram__00000006_.ProductComposition_Shapes = append(__Diagram__00000006_.ProductComposition_Shapes, __ProductCompositionShape__00000029_)
	__Diagram__00000006_.ProductComposition_Shapes = append(__Diagram__00000006_.ProductComposition_Shapes, __ProductCompositionShape__00000034_)
	__Diagram__00000006_.ProductComposition_Shapes = append(__Diagram__00000006_.ProductComposition_Shapes, __ProductCompositionShape__00000035_)
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
	__Library__00000000_.RootProducts = append(__Library__00000000_.RootProducts, __Product__00000007_)
	__Library__00000000_.RootProducts = append(__Library__00000000_.RootProducts, __Product__00000003_)
	__Library__00000000_.RootTasks = append(__Library__00000000_.RootTasks, __Task__00000001_)
	__Library__00000000_.RootTasks = append(__Library__00000000_.RootTasks, __Task__00000002_)
	__Library__00000000_.RootTasks = append(__Library__00000000_.RootTasks, __Task__00000003_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000003_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000006_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000000_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000001_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000002_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000003_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000004_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000005_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000006_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000007_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000008_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000009_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000010_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000011_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000012_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000013_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000014_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000015_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000016_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000017_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000018_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000019_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000020_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000021_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000022_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000023_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000024_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000025_)
	__Library__00000000_.Diagrams = append(__Library__00000000_.Diagrams, __Diagram__00000001_)
	__Library__00000000_.Diagrams = append(__Library__00000000_.Diagrams, __Diagram__00000002_)
	__Library__00000000_.Diagrams = append(__Library__00000000_.Diagrams, __Diagram__00000003_)
	__Library__00000000_.Diagrams = append(__Library__00000000_.Diagrams, __Diagram__00000004_)
	__Library__00000000_.Diagrams = append(__Library__00000000_.Diagrams, __Diagram__00000005_)
	__Library__00000000_.Diagrams = append(__Library__00000000_.Diagrams, __Diagram__00000006_)
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
	__Note__00000013_.Products = append(__Note__00000013_.Products, __Product__00000034_)
	__Note__00000014_.Products = append(__Note__00000014_.Products, __Product__00000034_)
	__Note__00000015_.Products = append(__Note__00000015_.Products, __Product__00000036_)
	__Note__00000016_.Products = append(__Note__00000016_.Products, __Product__00000039_)
	__Note__00000017_.Tasks = append(__Note__00000017_.Tasks, __Task__00000007_)
	__Note__00000018_.Products = append(__Note__00000018_.Products, __Product__00000039_)
	__Note__00000019_.Products = append(__Note__00000019_.Products, __Product__00000041_)
	__Note__00000020_.Products = append(__Note__00000020_.Products, __Product__00000042_)
	__Note__00000021_.Products = append(__Note__00000021_.Products, __Product__00000042_)
	__Note__00000022_.Products = append(__Note__00000022_.Products, __Product__00000044_)
	__Note__00000023_.Products = append(__Note__00000023_.Products, __Product__00000044_)
	__Note__00000024_.Products = append(__Note__00000024_.Products, __Product__00000044_)
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
	__NoteProductShape__00000016_.Note = __Note__00000013_
	__NoteProductShape__00000016_.Product = __Product__00000034_
	__NoteProductShape__00000017_.Note = __Note__00000014_
	__NoteProductShape__00000017_.Product = __Product__00000034_
	__NoteProductShape__00000018_.Note = __Note__00000015_
	__NoteProductShape__00000018_.Product = __Product__00000036_
	__NoteProductShape__00000019_.Note = __Note__00000018_
	__NoteProductShape__00000019_.Product = __Product__00000039_
	__NoteProductShape__00000020_.Note = __Note__00000016_
	__NoteProductShape__00000020_.Product = __Product__00000039_
	__NoteProductShape__00000021_.Note = __Note__00000019_
	__NoteProductShape__00000021_.Product = __Product__00000041_
	__NoteProductShape__00000022_.Note = __Note__00000020_
	__NoteProductShape__00000022_.Product = __Product__00000042_
	__NoteProductShape__00000023_.Note = __Note__00000021_
	__NoteProductShape__00000023_.Product = __Product__00000042_
	__NoteProductShape__00000024_.Note = __Note__00000022_
	__NoteProductShape__00000024_.Product = __Product__00000044_
	__NoteProductShape__00000025_.Note = __Note__00000023_
	__NoteProductShape__00000025_.Product = __Product__00000044_
	__NoteProductShape__00000026_.Note = __Note__00000024_
	__NoteProductShape__00000026_.Product = __Product__00000044_
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
	__NoteShape__00000013_.Note = __Note__00000013_
	__NoteShape__00000014_.Note = __Note__00000014_
	__NoteShape__00000015_.Note = __Note__00000015_
	__NoteShape__00000016_.Note = __Note__00000016_
	__NoteShape__00000017_.Note = __Note__00000017_
	__NoteShape__00000018_.Note = __Note__00000018_
	__NoteShape__00000019_.Note = __Note__00000019_
	__NoteShape__00000020_.Note = __Note__00000020_
	__NoteShape__00000021_.Note = __Note__00000021_
	__NoteShape__00000022_.Note = __Note__00000022_
	__NoteShape__00000023_.Note = __Note__00000023_
	__NoteShape__00000024_.Note = __Note__00000024_
	__NoteShape__00000025_.Note = __Note__00000025_
	__NoteTaskShape__00000001_.Note = __Note__00000017_
	__NoteTaskShape__00000001_.Task = __Task__00000007_
	__Product__00000001_.ReferencedProduct = nil
	__Product__00000002_.ReferencedProduct = nil
	__Product__00000002_.SubProducts = append(__Product__00000002_.SubProducts, __Product__00000011_)
	__Product__00000002_.SubProducts = append(__Product__00000002_.SubProducts, __Product__00000014_)
	__Product__00000002_.SubProducts = append(__Product__00000002_.SubProducts, __Product__00000019_)
	__Product__00000002_.SubProducts = append(__Product__00000002_.SubProducts, __Product__00000021_)
	__Product__00000002_.SubProducts = append(__Product__00000002_.SubProducts, __Product__00000025_)
	__Product__00000002_.SubProducts = append(__Product__00000002_.SubProducts, __Product__00000026_)
	__Product__00000002_.SubProducts = append(__Product__00000002_.SubProducts, __Product__00000033_)
	__Product__00000003_.ReferencedProduct = nil
	__Product__00000003_.SubProducts = append(__Product__00000003_.SubProducts, __Product__00000004_)
	__Product__00000003_.SubProducts = append(__Product__00000003_.SubProducts, __Product__00000005_)
	__Product__00000003_.SubProducts = append(__Product__00000003_.SubProducts, __Product__00000022_)
	__Product__00000004_.ReferencedProduct = nil
	__Product__00000005_.ReferencedProduct = nil
	__Product__00000005_.SubProducts = append(__Product__00000005_.SubProducts, __Product__00000017_)
	__Product__00000005_.SubProducts = append(__Product__00000005_.SubProducts, __Product__00000009_)
	__Product__00000005_.SubProducts = append(__Product__00000005_.SubProducts, __Product__00000029_)
	__Product__00000006_.ReferencedProduct = nil
	__Product__00000006_.SubProducts = append(__Product__00000006_.SubProducts, __Product__00000002_)
	__Product__00000006_.SubProducts = append(__Product__00000006_.SubProducts, __Product__00000001_)
	__Product__00000006_.SubProducts = append(__Product__00000006_.SubProducts, __Product__00000010_)
	__Product__00000007_.ReferencedProduct = nil
	__Product__00000007_.SubProducts = append(__Product__00000007_.SubProducts, __Product__00000006_)
	__Product__00000007_.SubProducts = append(__Product__00000007_.SubProducts, __Product__00000008_)
	__Product__00000008_.ReferencedProduct = nil
	__Product__00000009_.ReferencedProduct = nil
	__Product__00000009_.SubProducts = append(__Product__00000009_.SubProducts, __Product__00000023_)
	__Product__00000009_.SubProducts = append(__Product__00000009_.SubProducts, __Product__00000024_)
	__Product__00000009_.SubProducts = append(__Product__00000009_.SubProducts, __Product__00000027_)
	__Product__00000010_.ReferencedProduct = nil
	__Product__00000011_.ReferencedProduct = nil
	__Product__00000011_.SubProducts = append(__Product__00000011_.SubProducts, __Product__00000012_)
	__Product__00000012_.ReferencedProduct = nil
	__Product__00000012_.SubProducts = append(__Product__00000012_.SubProducts, __Product__00000013_)
	__Product__00000013_.ReferencedProduct = nil
	__Product__00000013_.SubProducts = append(__Product__00000013_.SubProducts, __Product__00000020_)
	__Product__00000013_.SubProducts = append(__Product__00000013_.SubProducts, __Product__00000028_)
	__Product__00000014_.ReferencedProduct = nil
	__Product__00000014_.SubProducts = append(__Product__00000014_.SubProducts, __Product__00000015_)
	__Product__00000015_.ReferencedProduct = nil
	__Product__00000015_.SubProducts = append(__Product__00000015_.SubProducts, __Product__00000016_)
	__Product__00000016_.ReferencedProduct = nil
	__Product__00000017_.ReferencedProduct = nil
	__Product__00000017_.SubProducts = append(__Product__00000017_.SubProducts, __Product__00000018_)
	__Product__00000018_.ReferencedProduct = nil
	__Product__00000019_.ReferencedProduct = nil
	__Product__00000020_.ReferencedProduct = nil
	__Product__00000021_.ReferencedProduct = nil
	__Product__00000022_.ReferencedProduct = nil
	__Product__00000023_.ReferencedProduct = nil
	__Product__00000024_.ReferencedProduct = nil
	__Product__00000025_.ReferencedProduct = nil
	__Product__00000026_.ReferencedProduct = nil
	__Product__00000027_.ReferencedProduct = nil
	__Product__00000028_.ReferencedProduct = nil
	__Product__00000029_.ReferencedProduct = nil
	__Product__00000029_.SubProducts = append(__Product__00000029_.SubProducts, __Product__00000030_)
	__Product__00000029_.SubProducts = append(__Product__00000029_.SubProducts, __Product__00000031_)
	__Product__00000029_.SubProducts = append(__Product__00000029_.SubProducts, __Product__00000032_)
	__Product__00000029_.SubProducts = append(__Product__00000029_.SubProducts, __Product__00000034_)
	__Product__00000029_.SubProducts = append(__Product__00000029_.SubProducts, __Product__00000035_)
	__Product__00000029_.SubProducts = append(__Product__00000029_.SubProducts, __Product__00000037_)
	__Product__00000030_.ReferencedProduct = nil
	__Product__00000031_.ReferencedProduct = nil
	__Product__00000032_.ReferencedProduct = nil
	__Product__00000033_.ReferencedProduct = nil
	__Product__00000033_.SubProducts = append(__Product__00000033_.SubProducts, __Product__00000038_)
	__Product__00000034_.ReferencedProduct = nil
	__Product__00000035_.ReferencedProduct = nil
	__Product__00000035_.SubProducts = append(__Product__00000035_.SubProducts, __Product__00000036_)
	__Product__00000035_.SubProducts = append(__Product__00000035_.SubProducts, __Product__00000039_)
	__Product__00000035_.SubProducts = append(__Product__00000035_.SubProducts, __Product__00000040_)
	__Product__00000035_.SubProducts = append(__Product__00000035_.SubProducts, __Product__00000041_)
	__Product__00000035_.SubProducts = append(__Product__00000035_.SubProducts, __Product__00000042_)
	__Product__00000035_.SubProducts = append(__Product__00000035_.SubProducts, __Product__00000043_)
	__Product__00000035_.SubProducts = append(__Product__00000035_.SubProducts, __Product__00000045_)
	__Product__00000036_.ReferencedProduct = nil
	__Product__00000037_.ReferencedProduct = nil
	__Product__00000038_.ReferencedProduct = nil
	__Product__00000039_.ReferencedProduct = nil
	__Product__00000040_.ReferencedProduct = nil
	__Product__00000041_.ReferencedProduct = nil
	__Product__00000042_.ReferencedProduct = nil
	__Product__00000043_.ReferencedProduct = nil
	__Product__00000043_.SubProducts = append(__Product__00000043_.SubProducts, __Product__00000044_)
	__Product__00000044_.ReferencedProduct = nil
	__Product__00000045_.ReferencedProduct = nil
	__Product__00000045_.SubProducts = append(__Product__00000045_.SubProducts, __Product__00000046_)
	__Product__00000046_.ReferencedProduct = nil
	__ProductCompositionShape__00000000_.Product = __Product__00000004_
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
	__ProductCompositionShape__00000016_.Product = __Product__00000020_
	__ProductCompositionShape__00000018_.Product = __Product__00000022_
	__ProductCompositionShape__00000019_.Product = __Product__00000023_
	__ProductCompositionShape__00000020_.Product = __Product__00000024_
	__ProductCompositionShape__00000021_.Product = __Product__00000025_
	__ProductCompositionShape__00000022_.Product = __Product__00000026_
	__ProductCompositionShape__00000023_.Product = __Product__00000027_
	__ProductCompositionShape__00000024_.Product = __Product__00000028_
	__ProductCompositionShape__00000025_.Product = __Product__00000029_
	__ProductCompositionShape__00000026_.Product = __Product__00000030_
	__ProductCompositionShape__00000027_.Product = __Product__00000031_
	__ProductCompositionShape__00000028_.Product = __Product__00000032_
	__ProductCompositionShape__00000029_.Product = __Product__00000033_
	__ProductCompositionShape__00000030_.Product = __Product__00000034_
	__ProductCompositionShape__00000031_.Product = __Product__00000035_
	__ProductCompositionShape__00000032_.Product = __Product__00000036_
	__ProductCompositionShape__00000033_.Product = __Product__00000037_
	__ProductCompositionShape__00000034_.Product = __Product__00000038_
	__ProductCompositionShape__00000035_.Product = __Product__00000021_
	__ProductCompositionShape__00000036_.Product = __Product__00000039_
	__ProductCompositionShape__00000037_.Product = __Product__00000040_
	__ProductCompositionShape__00000038_.Product = __Product__00000041_
	__ProductCompositionShape__00000039_.Product = __Product__00000042_
	__ProductCompositionShape__00000040_.Product = __Product__00000043_
	__ProductCompositionShape__00000041_.Product = __Product__00000044_
	__ProductCompositionShape__00000042_.Product = __Product__00000045_
	__ProductCompositionShape__00000043_.Product = __Product__00000046_
	__ProductCompositionShape__00000044_.Product = __Product__00000005_
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
	__ProductShape__00000027_.Product = __Product__00000022_
	__ProductShape__00000028_.Product = __Product__00000023_
	__ProductShape__00000029_.Product = __Product__00000024_
	__ProductShape__00000030_.Product = __Product__00000025_
	__ProductShape__00000031_.Product = __Product__00000026_
	__ProductShape__00000032_.Product = __Product__00000027_
	__ProductShape__00000033_.Product = __Product__00000028_
	__ProductShape__00000034_.Product = __Product__00000029_
	__ProductShape__00000035_.Product = __Product__00000030_
	__ProductShape__00000036_.Product = __Product__00000031_
	__ProductShape__00000037_.Product = __Product__00000032_
	__ProductShape__00000038_.Product = __Product__00000033_
	__ProductShape__00000039_.Product = __Product__00000034_
	__ProductShape__00000040_.Product = __Product__00000035_
	__ProductShape__00000041_.Product = __Product__00000036_
	__ProductShape__00000042_.Product = __Product__00000037_
	__ProductShape__00000043_.Product = __Product__00000038_
	__ProductShape__00000044_.Product = __Product__00000021_
	__ProductShape__00000045_.Product = __Product__00000039_
	__ProductShape__00000046_.Product = __Product__00000040_
	__ProductShape__00000047_.Product = __Product__00000041_
	__ProductShape__00000048_.Product = __Product__00000042_
	__ProductShape__00000049_.Product = __Product__00000043_
	__ProductShape__00000050_.Product = __Product__00000044_
	__ProductShape__00000051_.Product = __Product__00000045_
	__ProductShape__00000052_.Product = __Product__00000046_
	__Resource__00000000_.ReferencedResource = nil
	__Resource__00000000_.Tasks = append(__Resource__00000000_.Tasks, __Task__00000001_)
	__Resource__00000001_.ReferencedResource = nil
	__Resource__00000002_.ReferencedResource = nil
	__Resource__00000003_.ReferencedResource = nil
	__Resource__00000003_.SubResources = append(__Resource__00000003_.SubResources, __Resource__00000004_)
	__Resource__00000003_.SubResources = append(__Resource__00000003_.SubResources, __Resource__00000000_)
	__Resource__00000004_.ReferencedResource = nil
	__Resource__00000004_.SubResources = append(__Resource__00000004_.SubResources, __Resource__00000005_)
	__Resource__00000004_.SubResources = append(__Resource__00000004_.SubResources, __Resource__00000009_)
	__Resource__00000004_.SubResources = append(__Resource__00000004_.SubResources, __Resource__00000010_)
	__Resource__00000005_.ReferencedResource = nil
	__Resource__00000005_.SubResources = append(__Resource__00000005_.SubResources, __Resource__00000001_)
	__Resource__00000005_.SubResources = append(__Resource__00000005_.SubResources, __Resource__00000002_)
	__Resource__00000006_.ReferencedResource = nil
	__Resource__00000006_.SubResources = append(__Resource__00000006_.SubResources, __Resource__00000007_)
	__Resource__00000007_.ReferencedResource = nil
	__Resource__00000009_.ReferencedResource = nil
	__Resource__00000010_.ReferencedResource = nil
	__ResourceCompositionShape__00000000_.Resource = __Resource__00000004_
	__ResourceCompositionShape__00000001_.Resource = __Resource__00000005_
	__ResourceCompositionShape__00000002_.Resource = __Resource__00000001_
	__ResourceCompositionShape__00000003_.Resource = __Resource__00000002_
	__ResourceCompositionShape__00000004_.Resource = __Resource__00000000_
	__ResourceCompositionShape__00000005_.Resource = __Resource__00000007_
	__ResourceCompositionShape__00000007_.Resource = __Resource__00000009_
	__ResourceCompositionShape__00000008_.Resource = __Resource__00000010_
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
	__ResourceShape__00000012_.Resource = __Resource__00000009_
	__ResourceShape__00000013_.Resource = __Resource__00000000_
	__ResourceShape__00000014_.Resource = __Resource__00000000_
	__ResourceShape__00000015_.Resource = __Resource__00000010_
	__ResourceTaskShape__00000000_.Resource = __Resource__00000000_
	__ResourceTaskShape__00000000_.Task = __Task__00000001_
	__Task__00000000_.ReferencedTask = nil
	__Task__00000000_.SubTasks = append(__Task__00000000_.SubTasks, __Task__00000007_)
	__Task__00000001_.ReferencedTask = nil
	__Task__00000001_.SubTasks = append(__Task__00000001_.SubTasks, __Task__00000004_)
	__Task__00000002_.ReferencedTask = nil
	__Task__00000002_.SubTasks = append(__Task__00000002_.SubTasks, __Task__00000005_)
	__Task__00000002_.SubTasks = append(__Task__00000002_.SubTasks, __Task__00000006_)
	__Task__00000002_.SubTasks = append(__Task__00000002_.SubTasks, __Task__00000000_)
	__Task__00000003_.ReferencedTask = nil
	__Task__00000004_.ReferencedTask = nil
	__Task__00000004_.Inputs = append(__Task__00000004_.Inputs, __Product__00000004_)
	__Task__00000004_.Outputs = append(__Task__00000004_.Outputs, __Product__00000005_)
	__Task__00000005_.ReferencedTask = nil
	__Task__00000006_.ReferencedTask = nil
	__Task__00000007_.ReferencedTask = nil
	__TaskCompositionShape__00000001_.Task = __Task__00000000_
	__TaskCompositionShape__00000002_.Task = __Task__00000004_
	__TaskCompositionShape__00000003_.Task = __Task__00000005_
	__TaskCompositionShape__00000004_.Task = __Task__00000006_
	__TaskCompositionShape__00000005_.Task = __Task__00000007_
	__TaskInputShape__00000000_.Product = __Product__00000004_
	__TaskInputShape__00000000_.Task = __Task__00000004_
	__TaskOutputShape__00000000_.Task = __Task__00000004_
	__TaskOutputShape__00000000_.Product = __Product__00000005_
	__TaskShape__00000003_.Task = __Task__00000001_
	__TaskShape__00000004_.Task = __Task__00000002_
	__TaskShape__00000005_.Task = __Task__00000000_
	__TaskShape__00000006_.Task = __Task__00000003_
	__TaskShape__00000007_.Task = __Task__00000001_
	__TaskShape__00000008_.Task = __Task__00000004_
	__TaskShape__00000009_.Task = __Task__00000004_
	__TaskShape__00000010_.Task = __Task__00000005_
	__TaskShape__00000011_.Task = __Task__00000006_
	__TaskShape__00000012_.Task = __Task__00000007_
}
