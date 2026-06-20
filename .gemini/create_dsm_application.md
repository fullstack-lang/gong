# Creating a DSM Application in Gong

This document outlines the systematic steps and best practices for creating a domain-specific modeling (DSM) application utilizing Gong, as demonstrated by the transition and creation of `dsm/structure`.

## 1. Domain Modeling (Abstract & Concrete Types)
- **Abstract Types (`models_abstract.go` / `models.go`)**: 
  - Define the core business logic structs (e.g., `Library`, `System`, `Part`, `Link`).
  - Equip models with necessary slice relationships and UI state trackers (e.g., `IsExpanded`, `IsSubSystemsNodeExpanded`, `SubSystemsWhoseNodeIsExpanded`).
- **Concrete Shapes (`models_concrete.go`)**: 
  - Define visual shapes corresponding to your abstract types (e.g., `PartShape`, `LinkAssociationShape`).
  - Embed Gong UI fields like `RectShape`, `LinkShape`, and `ConcreteTypeFields`.
  - Implement getter/setter interfaces (e.g. `GetAbstractElement()`, `SetAbstractElement()`, `GetConcreteLayoutDirection()`) to map back to the Abstract forms.
- **Diagram Root (`diagram_structure.go`)**:
  - Define a root Diagram struct (e.g., `DiagramStructure`) that holds slices of both abstract inclusions (`Parts`, `Links`) and concrete placements (`Part_Shapes`, `Link_Shapes`).
  - Track activation states via `IsChecked`.

## 2. Gong Code Generation
- Execute the generator frequently: `gong generate --dsm --skipGoModCommands`.
- This auto-generates the `yyy_*` boilerplate files, configuring layout directions, generating `addCreateItemButton` generic button callbacks, SVG struct mappers, and probe forms based on your explicit structs.

## 3. UI Tree Layout (`stager_ux_tree.go`)
- Create modular tree building files (e.g. `stager_ux_tree_library.go`, `stager_ux_tree_system.go`) that traverse your hierarchy starting from `getRootLibrary()`.
- Use `tree.Node` structures to represent folders, fields, and objects.
- **Critical Callbacks**:
  - Wire UI interactive elements properly to mirror the backend state:
    - `OnIsExpandedChange`: Map to `stager.onIsExpandedChangeBool(...)` for folder memory.
    - `OnNameChange`: Map to `stager.onNameChange(element)` for renaming objects inside the tree.
    - `OnClick`: Map to `onNodeClicked(stager, element)` to open properties in the data probe.
    - `OnIsCheckedChanged`: To toggle active diagrams.
  - Utilize generated generic buttons: Bind `addCreateItemButton` to allow nested object instantiation directly from tree folders. Bind `addCreateItemShapeAndLinkButton` to instantiate visual objects alongside their abstract representations automatically.

## 4. SVG Mapping and Interactivity (`stager_ux_svg.go`)
- Map your domain `PartShape` and `LinkAssociationShape`s to generic `svg.Rect` and `svg.Link` instances.
- Define `onUpdateElementInDiagram` logic within `yyy_on_update_element_in_diagram.go` (if overwritten) to push DOM drag-and-drop / resize events from the SVG canvas back to the underlying `X`, `Y`, `Width`, and `Height` domain state.

## 5. Semantic Enforcement (`stager_enforce_semantic.go`)
- The backend requires a dedicated multi-pass semantic loop that settles the domain state when modifications occur.
- **Standard rules include**:
  - `enforceShapeOrphans()`: Cleans up and unstages shapes if their underlying abstract elements are deleted.
  - `enforceShapesAbstractConsistency()`: Removes orphaned slices.
  - `enforceShapeNames()`: Crucial DSM practice. Enforces that `PartShape`s strictly inherit their names (e.g., `AbstractElementName-DiagramName`). This syncs the visual identity automatically when an abstract object is renamed in the UI tree.
  - Architectural constraints: Examples include `enforceAtLeastOneDiagramPerSystem()`, defaulting required relations to maintain application stability.

## Summary
The pipeline generally flows as: Model abstract logic -> Model UI layout properties -> Model visual shapes -> Run code-gen -> Build the visual SVG layout -> Wire the UX tree -> Add strict semantic checks to lock application integrity.
