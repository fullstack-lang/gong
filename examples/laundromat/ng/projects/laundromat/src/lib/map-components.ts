// insertion point sub template for components imports 
import { MachinesTableComponent } from './machines-table/machines-table.component'
import { SimulationsTableComponent } from './simulations-table/simulations-table.component'
import { WashersTableComponent } from './washers-table/washers-table.component'

// insertion point sub template for map of components per struct 
export const MapOfMachinesComponents: Map<string, any> = new Map([["MachinesTableComponent", MachinesTableComponent],])
export const MapOfSimulationsComponents: Map<string, any> = new Map([["SimulationsTableComponent", SimulationsTableComponent],])
export const MapOfWashersComponents: Map<string, any> = new Map([["WashersTableComponent", WashersTableComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Machine", MapOfMachinesComponents],
      ["Simulation", MapOfSimulationsComponents],
      ["Washer", MapOfWashersComponents],
    ]
  )