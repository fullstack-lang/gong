package main

import (
	"flag"
	"log"
	"strconv"
	"time"

	// insertion point for models import
	test3_level1stack "github.com/fullstack-lang/gong/test/test3/go/level1stack"
	test3_models "github.com/fullstack-lang/gong/test/test3/go/models"
	"github.com/fullstack-lang/gong/test/test3/go/probe"

	// static

	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")
)

func main() {

	log.SetPrefix("test3: ")
	log.SetFlags(log.Lmicroseconds)

	// parse program arguments
	flag.Parse()

	// setup model stack with its probe
	stack := test3_level1stack.NewLevel1Stack("test3", *unmarshallFromCode, *marshallOnCommit, true, *embeddedDiagrams)
	stack.Probe.Refresh()

	// the root split name is "" by convention. Is is the same for all gong applications
	// that do not develop their specific angular component
	splitStage := split_stack.NewStack(stack.R, "", "", "", "", false, false).Stage

	test3_models.NewStager(stack.R, stack.Stage, splitStage)

	stack.Probe.AddNotification(probe.Notification{
		Date: time.Now(),
		Message: `La rue assourdissante autour de moi hurlait.
Longue, mince, en grand deuil, douleur majestueuse,
Une femme passa, d’une main fastueuse
Soulevant, balançant le feston et l’ourlet;

Agile et noble, avec sa jambe de statue.
Moi, je buvais, crispé comme un extravagant,
Dans son oeil, ciel livide où germe l’ouragan,
La douceur qui fascine et le plaisir qui tue.

Un éclair… puis la nuit! – Fugitive beauté
Dont le regard m’a fait soudainement renaître,
Ne te verrai-je plus que dans l’éternité?

Ailleurs, bien loin d’ici! trop tard! jamais peut-être!
Car j’ignore où tu fuis, tu ne sais où je vais,
O toi que j’eusse aimée, ô toi qui le savais!`,
	})

	log.Println("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := stack.R.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
