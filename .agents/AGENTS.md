# call gong generate when ...

docs/gong-go-api.md describes the gong syntax

When you update an element that is part of this syntax, run `gong generate` in its `models` directory.

**CRITICAL RULE FOR RUNNING GONG GENERATE:**
You MUST run `gong generate` (or `gong generate --dsm`) with your Current Working Directory (`Cwd`) strictly set to the absolute path of the `go/models` directory of the project. 

- **DO NOT** run it from the root of the project.
- **DO NOT** run it from the `go` directory.
- **DO NOT** run `gong generate <path>`. Always set `Cwd` to the `go/models` directory, and run `gong generate` with no path arguments.

Failure to follow this rule will cause Gong to generate thousands of unwanted files in the wrong directories and break the workspace.

---

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

  note that the diagram node is a root of all elements of the library. for each element node below the diagram, check/uncheck the diagram add/remove the shape to the diagram 

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

---

## dsm projects are in the dsm directory

you should  use "gong generate --dsm" in the models directory with those if you want to gong generate

---

# Modifying DSM yyy files

When you are asked to modify any `yyy_*.go` files within the `dsm` directory, you **must adhere to the following rules**:

1. **Master Directory**: The `dsm/project` directory is the **master** source of truth for all `yyy_*.go` files.
2. **Do Not Edit Other Projects Directly**: Do NOT manually edit `yyy_*.go` files in any other `dsm` sub-projects (such as `process`, `barrgraph`, `statemachines`, etc.).
3. **Propagation Strategy**: 
   - Apply your modifications **only** to the `yyy_*.go` files located in `dsm/project/go/models/`.
   - Once your edits are complete in the master directory, run the command `gong generate --dsm` in the models directory of each dsm.
    - *never* run "`gong generate --dsm` in dsm/project/go/models because it will overwrite the changes to yyy files (you need to recompile the gong command to take into account the updated yyy files)


use the task
                "label": "go install gong & gong generate --dsm go/models in all dsm projects",

to propagate your change to all dsm

---

# Enforcing Shape Orphans in DSM

When adding or modifying shapes in a Domain Specific Model (DSM), you must ensure that all shapes are checked for being orphaned.

If a new shape type is added, you **must adhere to the following rules**:

1. **Update `enforceShapeOrphans`**: The new shape type must be accounted for in the `enforceShapeOrphans` method located in `dsm/<dsm name>/go/models/stager_enforce_shape_orphans.go`.
2. **Collect Reachable Shapes**: Inside the `for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage)` loop, collect all reachable instances of your new shape type from the diagram (e.g. `collectShapes(diagram.MyNewShapes, reachableMyNewShapes)`).
3. **Unstage Unreachable Shapes**: After the loop, use `unstageUnreachableOrphans` to remove any instances of your new shape type that are not attached to a diagram (e.g. `needCommit = unstageUnreachableOrphans(stager, reachableMyNewShapes) || needCommit`).

---

I want you to never commit anything to the gong repository.
ask me if you need to

---

# Gong Framework Rule: Avoid Successive Stack Updates
Rule: Do not invoke successive UI/Stage updates (such as calling stager.ux_tree() immediately after stager.stage.Commit()) within the same callback.

## Context
In the Gong framework, when an event happens (like an OnClick callback), developers often commit changes to update the backend database and push changes to the frontend.

## The Problem
Calling stager.stage.Commit() automatically triggers a regeneration process for related states (such as rebuilding the UI tree representation via stager_tree()).

At the commit, with a stack of heigth 4 like tree, the commit is captured by the orm and the commit number is sent to the stack controller websocket that multiplex flows from all tree stack instances. The controller will then copy the orm copy of the stage and send it to the front with the websocket.

If you immediately call another stack upda

This crash happens because the first commit is already tearing down and rebuilding tree elements (like Button nodes), and the second update attempts to reference or mutate those elements while they are in an inconsistent state or no longer exist.

This lack of robustness to successive commits is a bug but it is also a way to spot spurious calls to stage.Commit().


## Best Practice
Use a single update trigger: Inside a callback, only use a single state commitment call like stager.stage.Commit().
Rely on built-in regeneration: Trust that Commit() handles the propagation of the state correctly without manually forcing the tree or SVG to redraw a second time.
Separate concerns: If you absolutely need complex layout updates followed by a UI refresh, ensure they are handled within a single transactional update path or use suspended callbacks (CommitWithSuspendedCallbacks()) where appropriate.

---

dont forget the stroke opacity

for instance

		line2.Presentation.Stroke = "black"
		line2.Presentation.StrokeWidth = 1.0
		line2.Presentation.StrokeOpacity = 1.0

## taking into account a modification in a stack of the lib

if a stack in the lib, with height 4, is modified.

you need to gong generate lib/split/go/models for the height 0 stack to be affected. Indeed, height 0 stack import 

lib/split/ng-github.com-fullstack-lang-gong-lib-split/dist/ng-github.com-fullstack-lang-gong-lib-split

## Allowed path

You are always allowed to access "Download" elements

---

# Generating code for DSM projects

When generating code for DSM projects (which reside in the `dsm` directory), **never use the `--stack-height=4` flag**.

DSMs are stacks of height 0 and do not incorporate specific Angular component libraries directly. Using `--stack-height=4` on a DSM project will incorrectly scaffold a full Angular frontend application and create thousands of unwanted files.

**Correct Command for DSM projects:**
```bash
gong generate --dsm
```
(You may append `--skipGoModCommands` if necessary, but **NEVER** `--stack-height=4`).