package lasagna

func PreparationTime(layers []string, avg int) int {
	if avg == 0 {
		avg = 2
	}
	return len(layers) * avg
}

func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0

	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		} else if layer == "sauce" {
			sauce += 0.2
		}
	}

	return noodles, sauce
}

func AddSecretIngredient(friendIngredients, myIngredients []string) {
	if len(friendIngredients) == 0 || len(myIngredients) == 0 {
		return
	}
	if myIngredients[len(myIngredients)-1] == "?" {
		myIngredients[len(myIngredients)-1] = friendIngredients[len(friendIngredients)-1]
	}
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	if portions <= 0 {
		return nil
	}

	scaledQuantities := make([]float64, len(quantities))
	for i, quantity := range quantities {
		scaledQuantities[i] = quantity * float64(portions) / 2.0
	}
	return scaledQuantities
}
