package tests

import (
	"log"
	"testing"

	"github.com/fullstack-lang/gong/stacks/gong/go/models"
)

func TestPrimeNumber(t *testing.T) {

	// let's print them
	for _, x := range models.Primes {
		log.Println(x)
	}
}
