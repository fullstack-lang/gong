import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, combineLatest } from 'rxjs';

// insertion point sub template for services imports 
import { MachineDB } from './machine-db'
import { MachineService } from './machine.service'

import { SimulationDB } from './simulation-db'
import { SimulationService } from './simulation.service'

import { WasherDB } from './washer-db'
import { WasherService } from './washer.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template 
  Machines_array = new Array<MachineDB>(); // array of repo instances
  Machines = new Map<number, MachineDB>(); // map of repo instances
  Machines_batch = new Map<number, MachineDB>(); // same but only in last GET (for finding repo instances to delete)
  Simulations_array = new Array<SimulationDB>(); // array of repo instances
  Simulations = new Map<number, SimulationDB>(); // map of repo instances
  Simulations_batch = new Map<number, SimulationDB>(); // same but only in last GET (for finding repo instances to delete)
  Washers_array = new Array<WasherDB>(); // array of repo instances
  Washers = new Map<number, WasherDB>(); // map of repo instances
  Washers_batch = new Map<number, WasherDB>(); // same but only in last GET (for finding repo instances to delete)
}

//
// Store of all instances of the stack
//
export const FrontRepoSingloton = new (FrontRepo)

// define the type of nullable Int64 in order to support back pointers IDs
export class NullInt64 {
    Int64: number
    Valid: boolean
}

// define the interface for information that is forwarded from the calling instance to 
// the select table
export interface DialogData {
  ID: number; // ID of the calling instance
  ReversePointer: string; // field of {{Structname}} that serve as reverse pointer
}

//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
@Injectable({
  providedIn: 'root'
})
export class FrontRepoService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  constructor(
    private http: HttpClient, // insertion point sub template 
    private machineService: MachineService,
    private simulationService: SimulationService,
    private washerService: WasherService,
  ) { }

  // typing of observable can be messy in typescript. Therefore, one force the type
  observableFrontRepo: [ // insertion point sub template 
    Observable<MachineDB[]>,
    Observable<SimulationDB[]>,
    Observable<WasherDB[]>,
  ] = [ // insertion point sub template 
      this.machineService.getMachines(),
      this.simulationService.getSimulations(),
      this.washerService.getWashers(),
    ];

  //
  // pull performs a GET on all struct of the stack and redeem association pointers 
  //
  // This is an observable. Therefore, the control flow forks with
  // - pull() return immediatly the observable
  // - the observable observer, if it subscribe, is called when all GET calls are performs
  pull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest(
          this.observableFrontRepo
        ).subscribe(
          ([ // insertion point sub template for declarations 
            machines_,
            simulations_,
            washers_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var machines: MachineDB[]
            machines = machines_
            var simulations: SimulationDB[]
            simulations = simulations_
            var washers: WasherDB[]
            washers = washers_

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            FrontRepoSingloton.Machines_array = machines

            // clear the map that counts Machine in the GET
            FrontRepoSingloton.Machines_batch.clear()
            
            machines.forEach(
              machine => {
                FrontRepoSingloton.Machines.set(machine.ID, machine)
                FrontRepoSingloton.Machines_batch.set(machine.ID, machine)
              }
            )
            
            // clear machines that are absent from the batch
            FrontRepoSingloton.Machines.forEach(
              machine => {
                if (FrontRepoSingloton.Machines_batch.get(machine.ID) == undefined) {
                  FrontRepoSingloton.Machines.delete(machine.ID)
                }
              }
            )
            // init the array
            FrontRepoSingloton.Simulations_array = simulations

            // clear the map that counts Simulation in the GET
            FrontRepoSingloton.Simulations_batch.clear()
            
            simulations.forEach(
              simulation => {
                FrontRepoSingloton.Simulations.set(simulation.ID, simulation)
                FrontRepoSingloton.Simulations_batch.set(simulation.ID, simulation)
              }
            )
            
            // clear simulations that are absent from the batch
            FrontRepoSingloton.Simulations.forEach(
              simulation => {
                if (FrontRepoSingloton.Simulations_batch.get(simulation.ID) == undefined) {
                  FrontRepoSingloton.Simulations.delete(simulation.ID)
                }
              }
            )
            // init the array
            FrontRepoSingloton.Washers_array = washers

            // clear the map that counts Washer in the GET
            FrontRepoSingloton.Washers_batch.clear()
            
            washers.forEach(
              washer => {
                FrontRepoSingloton.Washers.set(washer.ID, washer)
                FrontRepoSingloton.Washers_batch.set(washer.ID, washer)
              }
            )
            
            // clear washers that are absent from the batch
            FrontRepoSingloton.Washers.forEach(
              washer => {
                if (FrontRepoSingloton.Washers_batch.get(washer.ID) == undefined) {
                  FrontRepoSingloton.Washers.delete(washer.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template for redeem 
            machines.forEach(
              machine => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            simulations.forEach(
              simulation => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Machine redeeming
                {
                  let _machine = FrontRepoSingloton.Machines.get(simulation.MachineID.Int64)
                  if (_machine) {
                    simulation.Machine = _machine
                  }
                }
                // insertion point for pointer field Washer redeeming
                {
                  let _washer = FrontRepoSingloton.Washers.get(simulation.WasherID.Int64)
                  if (_washer) {
                    simulation.Washer = _washer
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )
            washers.forEach(
              washer => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Machine redeeming
                {
                  let _machine = FrontRepoSingloton.Machines.get(washer.MachineID.Int64)
                  if (_machine) {
                    washer.Machine = _machine
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // insertion point for pull per struct 

  // MachinePull performs a GET on Machine of the stack and redeem association pointers 
  MachinePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.machineService.getMachines()
        ]).subscribe(
          ([ // insertion point sub template 
            machines,
          ]) => {
            // init the array
            FrontRepoSingloton.Machines_array = machines

            // clear the map that counts Machine in the GET
            FrontRepoSingloton.Machines_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            machines.forEach(
              machine => {
                FrontRepoSingloton.Machines.set(machine.ID, machine)
                FrontRepoSingloton.Machines_batch.set(machine.ID, machine)

                // insertion point for redeeming ONE/ZERO-ONE associations 

                // insertion point for redeeming ONE-MANY associations 
              }
            )

            // clear machines that are absent from the GET
            FrontRepoSingloton.Machines.forEach(
              machine => {
                if (FrontRepoSingloton.Machines_batch.get(machine.ID) == undefined) {
                  FrontRepoSingloton.Machines.delete(machine.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // SimulationPull performs a GET on Simulation of the stack and redeem association pointers 
  SimulationPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.simulationService.getSimulations()
        ]).subscribe(
          ([ // insertion point sub template 
            simulations,
          ]) => {
            // init the array
            FrontRepoSingloton.Simulations_array = simulations

            // clear the map that counts Simulation in the GET
            FrontRepoSingloton.Simulations_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            simulations.forEach(
              simulation => {
                FrontRepoSingloton.Simulations.set(simulation.ID, simulation)
                FrontRepoSingloton.Simulations_batch.set(simulation.ID, simulation)

                // insertion point for redeeming ONE/ZERO-ONE associations 
                // insertion point for pointer field Machine redeeming
                {
                  let _machine = FrontRepoSingloton.Machines.get(simulation.MachineID.Int64)
                  if (_machine) {
                    simulation.Machine = _machine
                  }
                }
                // insertion point for pointer field Washer redeeming
                {
                  let _washer = FrontRepoSingloton.Washers.get(simulation.WasherID.Int64)
                  if (_washer) {
                    simulation.Washer = _washer
                  }
                }

                // insertion point for redeeming ONE-MANY associations 
              }
            )

            // clear simulations that are absent from the GET
            FrontRepoSingloton.Simulations.forEach(
              simulation => {
                if (FrontRepoSingloton.Simulations_batch.get(simulation.ID) == undefined) {
                  FrontRepoSingloton.Simulations.delete(simulation.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // WasherPull performs a GET on Washer of the stack and redeem association pointers 
  WasherPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.washerService.getWashers()
        ]).subscribe(
          ([ // insertion point sub template 
            washers,
          ]) => {
            // init the array
            FrontRepoSingloton.Washers_array = washers

            // clear the map that counts Washer in the GET
            FrontRepoSingloton.Washers_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            washers.forEach(
              washer => {
                FrontRepoSingloton.Washers.set(washer.ID, washer)
                FrontRepoSingloton.Washers_batch.set(washer.ID, washer)

                // insertion point for redeeming ONE/ZERO-ONE associations 
                // insertion point for pointer field Machine redeeming
                {
                  let _machine = FrontRepoSingloton.Machines.get(washer.MachineID.Int64)
                  if (_machine) {
                    washer.Machine = _machine
                  }
                }

                // insertion point for redeeming ONE-MANY associations 
              }
            )

            // clear washers that are absent from the GET
            FrontRepoSingloton.Washers.forEach(
              washer => {
                if (FrontRepoSingloton.Washers_batch.get(washer.ID) == undefined) {
                  FrontRepoSingloton.Washers.delete(washer.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }
}

// insertion point for get unique ID per struct 
export function getMachineUniqueID(id: number): number {
  return 31 * id
}
export function getSimulationUniqueID(id: number): number {
  return 37 * id
}
export function getWasherUniqueID(id: number): number {
  return 41 * id
}
