## Situation

Currently, the API

```go
OnAfterUpdate[Type]( stage * StageStruct, old, new Type)
```

is widely used.

### in probes

something similar (whaou it is widely confusiong because there is the generic API 
and the non generic API, this is a nightmare !)

```go
func (nodeImplGongstruct *TreeNodeImplGongstruct) OnAfterUpdate(
	gongtreeStage *gongtree_models.Stage,
	stagedNode, frontNode *gongtree_models.Node) {

	// setting the value of the staged node	to the new value
	// otherwise, the expansion memory is lost
	if stagedNode.IsExpanded != frontNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded
		return
	}
```

### Code generation if a gongstruct has the function signature

go/models/check_function_signature.go

```go
func checkFunctionSignature(file *ast.File, modelPkg *ModelPkg) {
...
```

```go
	// HasOnAfterUpdateSignature is used to generate orchestrator code
	HasOnAfterUpdateSignature bool
```

Therefore, in *this case* (the orchestrator), the API signature is stupid.
In the *other case*, the one in the probe, the API signature makes sense because the AfterUpdate is different
if the stageNode has a certain type or not.

```go
// OnAfterUpdate, notice that node == stagedNode
func (node *Node) OnAfterUpdate(stage *Stage, _, frontNode *Node) {

	if node.Impl != nil {
		node.Impl.OnAfterUpdate(stage, node, frontNode)
	}
}
```
Very confusing indeed.

```go
type NodeImplInterface interface {

	// OnAfterUpdate function is called each time a node is modified
	OnAfterUpdate(stage *Stage, stagedNode, frontNode *Node)
}

```


## Problems


- not the idiomatic way
- the `old` argument is sometimes not used if it is the receiver
- the **mouse event** is not forwarded and it is the philosophy of gong to forward all that is front to the back

### Not the idiomatic way

Here is was gemini says

Name potentially confusing. The ambiguity often comes from what "Update" and "Commit" refer to in the context of a REST API call and a database transaction.

#### Execution Flow
To clarify the timing, here is the sequence of events within a typical Update controller function in your code:

- A PATCH request is received for a resource.
- The corresponding record is fetched from the database (e.g., astructDB).
- The fields of this database object are updated from the request body.
- The changes are persisted to the database using db.Updates(&astructDB). This is the database write operation.
- A new model instance (e.g., astructNew) is created from the updated database object.
- The old model instance (astructOld) is retrieved from the in-memory "stage".
- The models.AfterUpdateFromFront(backRepo.GetStage(), astructOld, astructNew) function is called.
- The backRepo.IncrementPushFromFrontNb() function is called, which appears to be for state synchronization rather than a database commit.
- An HTTP 200 OK response is sent back to the client.

So, your understanding is correct that the callback happens during the REST API call. However, it happens after the specific database record has been updated. The confusion might arise from thinking of the entire API call as a single "update" transaction, or from the term "commit to the back" which seems to refer to something other than a database transaction commit.

#### Idiomatic Naming and Suggestions

The name AfterUpdateFromFront is descriptive, but as you've noted, it can be ambiguous. A more idiomatic approach in many frameworks for hooks, events, and callbacks is to use the On<Event> prefix. This convention is clear and widely understood.

Furthermore, the full callback chain (OnAfterAstructUpdateCallback.OnAfterUpdate) is a bit verbose.

I suggest a refactoring to a more concise and conventional naming scheme:

- Dispatcher functions: After<Action>FromFront → On<Action>FromFront (e.g., AfterUpdateFromFront → OnUpdateFromFront).
- Interfaces: OnAfter<Action>Interface → On<Action>Interface (e.g., OnAfterUpdateInterface → OnUpdateInterface).
- Interface methods: OnAfter<Action> → On<Action> (e.g., OnAfterUpdate → OnUpdate).
- Stage callback fields: OnAfter<Struct><Action>Callback → On<Struct><Action>Callback (e.g., OnAfterAstructUpdateCallback → OnAstructUpdateCallback).

This makes the naming more streamlined and aligned with common practices. Here are the code changes to implement this suggestion across the test stack.

## more use cases

There are use cases when one needs to know if the CTRL or SHIFT key is pressed.


## Solutions


### The make it alive both API

- depreciate the as-is API and writes a brand new API all the related go code for the mouse event (if len of params > 2)
- 

### Hard refactoring

- rewrite all API and let the code be refactored ()

<TO DO>

in level 2 applications, how many usages ?
