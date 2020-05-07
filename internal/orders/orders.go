package orders

import (
	"github.com/j-boivive/autochef/internal/constants"
	"math/rand"
)

func GenerateVerb(rand *rand.Rand) string {
	index := rand.Intn(len(constants.Verbs))
	return constants.Verbs[index]
}

func GenerateSubj(rand *rand.Rand) string {
	index := rand.Intn(len(constants.Subjs))
	return constants.Subjs[index]
}
