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

has 2 problems, 

- the `old` argument is sometimes not used if it is the receiver
- the **mouse event** is not forwarded and it is the philosophy of gong to forward all that is front to the back

There are use cases when one needs to know if the CTRL or SHIFT key is pressed.

Alternative solution:

- depreciate the as-is API and rewrites all the related go code
- 

<TO DO>

in level 2 applications, how many usages ?
