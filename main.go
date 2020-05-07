package main

import (
	"fmt"
	"github.com/j-boivive/autochef/internal/orders"
	"math/rand"
	"strings"
	"time"
)

//getRand returns a pseudorandom generator seeded with current time.
func getRand() *rand.Rand {
	s := rand.NewSource(time.Now().UnixNano())
	return rand.New(s)
}

//giveOrder makes autochef tell you what to do.
func giveOrder() {
	r := getRand()
	verb := orders.GenerateVerb(r)
	subj := orders.GenerateSubj(r)
	fmt.Printf("Chefen säger att du ska %s %s!\n", strings.ToLower(verb), subj)
}

func main() {
	giveOrder()
}
