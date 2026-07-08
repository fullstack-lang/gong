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