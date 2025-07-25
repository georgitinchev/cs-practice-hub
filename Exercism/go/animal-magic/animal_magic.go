package chance

import (
	"math/rand"
)

func RollADie() int {
	// RollADie returns a random int d with 1 <= d <= 20.
	n := rand.Intn(20)
	if n == 0 {
		return 1
	} else {
		return n
	}
}

func GenerateWandEnergy() float64 {
	// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
	return rand.Float64() * 12.0
}

func ShuffleAnimals() []string {
	// ShuffleAnimals returns a slice with all eight animal strings in random order.
	// animals are - ant,beaver,cat,dog,elephant,fox,giraffe,hedgehog
	animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	rand.Shuffle(len(animals), func(i, j int) {
		animals[i], animals[j] = animals[j], animals[i]
	})
	return animals
}
