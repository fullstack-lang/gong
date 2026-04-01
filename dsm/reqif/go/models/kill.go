package models

import "os"

// Kill is a gong struct that is used to kill the application by a remote call
// with curl
/*

```bash
curl --request POST \
  --url http://localhost:8080/api/github.com/fullstack-lang/gong/dsm/reqif/go/v1/kills?Name="reqif" \
  --header 'content-type: application/json' \
  --data '{"Name": "Foo"}'
```
*/
type Kill struct {
	Name string
}

type OnAfterKillCreateCallback struct {
}

// OnAfterCreate implements OnAfterCreateInterface.
func (o *OnAfterKillCreateCallback) OnAfterCreate(stage *Stage, instance *Kill) {
	os.Exit(0)
}
