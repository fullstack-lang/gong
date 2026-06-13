# Modifying DSM yyy files

When you are asked to modify any `yyy_*.go` files within the `dsm` directory, you **must adhere to the following rules**:

1. **Master Directory**: The `dsm/project` directory is the **master** source of truth for all `yyy_*.go` files.
2. **Do Not Edit Other Projects Directly**: Do NOT manually edit `yyy_*.go` files in any other `dsm` sub-projects (such as `process`, `barrgraph`, `statemachines`, etc.).
3. **Propagation Strategy**: 
   - Apply your modifications **only** to the `yyy_*.go` files located in `dsm/project/go/models/`.
   - Once your edits are complete in the master directory, run the command `gong generate --dsm` from the root of the repository.
   - This command will automatically propagate your updates to the corresponding `yyy_*.go` files in all other `dsm` projects.
