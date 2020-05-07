package orders

import (
    "github.com/j-boivive/autochef/internal/constants"
    "math/rand"
    "testing"
)

func getRand() *rand.Rand {
    s := rand.NewSource(1)
    return rand.New(s)
}
func runTest(t *testing.T, actual string, expected []string) {
    for _, word := range expected {
        if actual == word {
            return
        }
    }
    t.Errorf("%s, not in %s", actual, expected)
}

func TestGenerateVerb(t *testing.T) {
    r := getRand()
    v := GenerateVerb(r)
    runTest(t, v, constants.Verbs)
}

func TestGenerateSubj(t *testing.T) {
    r := getRand()
    s := GenerateSubj(r)
    runTest(t, s, constants.Subjs)
}
