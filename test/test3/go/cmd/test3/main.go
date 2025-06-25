package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"

	// insertion point for models import
	test3_models "github.com/fullstack-lang/gong/test/test3/go/models"
	test3_stack "github.com/fullstack-lang/gong/test/test3/go/stack"
	test3_static "github.com/fullstack-lang/gong/test/test3/go/static"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")
)

const nbInstances = 5000

func main() {

	log.SetPrefix("test3: ")
	log.SetFlags(log.Lmicroseconds)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := test3_static.ServeStaticFiles(*logGINFlag)

	// setup model stack with its probe
	log.Println("Before new Stack")
	start := time.Now()
	stack := test3_stack.NewStack(r, "test3", *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	stack.Probe.Refresh()
	duration := time.Since(start).Milliseconds()
	log.Printf("%d, per int %f", duration, float64(duration)/nbInstances)

	// the root split name is "" by convention. Is is the same for all gong applications
	// that do not develop their specific angular component

	splitStage := split_stack.NewStack(r, "", "", "", "", false, false).Stage
	log.Println("After new Stack")

	stager := test3_models.NewStager(r, stack.Stage, splitStage)

	as := *test3_models.GetGongstructInstancesSet[test3_models.A](stack.Stage)

	if *unmarshallFromCode == "data/stage-large.go" && len(as) == 0 {
		for i := range nbInstances {
			(&test3_models.A{Name: fmt.Sprintf("%.5d", i)}).Stage(stack.Stage)
		}
		stack.Stage.Commit()
	}

	// one for the probe of the
	split.StageBranch(splitStage, &split.View{
		Name: stack.Stage.GetName() + "with Probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Size: 50,
				AsSplit: (&split.AsSplit{
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						stager.GetAsSplitArea(),
					},
				}),
			}),
			(&split.AsSplitArea{
				Size: 50,
				Split: (&split.Split{
					StackName: stack.Stage.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	// commit the split stage (this will initiate the front components)
	splitStage.Commit()

	log.Println("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}

// countLinesBySymbol counts the occurrences of a specific line pattern in a file
// and aggregates the counts by a captured symbol within that pattern.
// The target pattern is: __A__000000_00000 := (&models.A{}).Stage(stage)
// where 'A' is the symbol to be captured.
func countLinesBySymbol(filename string) (map[string]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	counts := make(map[string]int)
	// Regex to match the line and capture the symbol (e.g., "A")
	re := regexp.MustCompile(`__([A-Z]+)__\d{6}_\d{5} := \(&models\.([A-Z]+)\{\}\)\.Stage\(stage\)`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			// The first submatch (index 1) is the captured symbol
			symbol := matches[1]
			counts[symbol]++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning file: %w", err)
	}

	return counts, nil
}
