package commands

import (
	"log"
	"strconv"

	test2_stack "github.com/fullstack-lang/gong/test/test2/go/stack"
	test2_static "github.com/fullstack-lang/gong/test/test2/go/static"
)

func executeServer(unmarshallFromCode string, marshallOnCommit string, port int, embeddedDiagrams bool, logGINFlag bool) {

	// setup the static file server and get the controller
	r := test2_static.ServeStaticFiles(logGINFlag)

	// setup stack
	stack := test2_stack.NewStack(r, "test2", unmarshallFromCode, marshallOnCommit, "", embeddedDiagrams, true)
	stack.Probe.Refresh()

	log.Println("Server ready serve on localhost:" + strconv.Itoa(port))
	err := r.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
