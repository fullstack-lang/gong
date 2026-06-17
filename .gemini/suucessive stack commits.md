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