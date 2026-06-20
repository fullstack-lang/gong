
# Enforcing Shape Orphans in DSM

When adding or modifying shapes in a Domain Specific Model (DSM), you must ensure that all shapes are checked for being orphaned.

If a new shape type is added, you **must adhere to the following rules**:

1. **Update `enforceShapeOrphans`**: The new shape type must be accounted for in the `enforceShapeOrphans` method located in `dsm/<dsm name>/go/models/stager_enforce_shape_orphans.go`.
2. **Collect Reachable Shapes**: Inside the `for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage)` loop, collect all reachable instances of your new shape type from the diagram (e.g. `collectShapes(diagram.MyNewShapes, reachableMyNewShapes)`).
3. **Unstage Unreachable Shapes**: After the loop, use `unstageUnreachableOrphans` to remove any instances of your new shape type that are not attached to a diagram (e.g. `needCommit = unstageUnreachableOrphans(stager, reachableMyNewShapes) || needCommit`).