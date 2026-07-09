# call gong generate when ...

docs/gong-go-api.md describes the gong syntax

When you update an element that is part of this syntax, run `gong generate` in its `models` directory.

**CRITICAL RULE FOR RUNNING GONG GENERATE:**
You MUST run `gong generate` (or `gong generate --dsm`) with your Current Working Directory (`Cwd`) strictly set to the absolute path of the `go/models` directory of the project. 

- **DO NOT** run it from the root of the project.
- **DO NOT** run it from the `go` directory.
- **DO NOT** run `gong generate <path>`. Always set `Cwd` to the `go/models` directory, and run `gong generate` with no path arguments.

Failure to follow this rule will cause Gong to generate thousands of unwanted files in the wrong directories and break the workspace.
