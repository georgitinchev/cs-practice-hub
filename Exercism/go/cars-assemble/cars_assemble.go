package cars

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	productionRatePerHour := float64(productionRate)
	successRateDecimal := successRate / 100.0
	workingCarsPerHour := productionRatePerHour * successRateDecimal
	return workingCarsPerHour
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	workingCarsPerHour := CalculateWorkingCarsPerHour(productionRate, successRate)
	workingCarsPerMinute := int(workingCarsPerHour / 60.0)
	return workingCarsPerMinute
}

func CalculateCost(carsCount int) uint {
	const costPerTenCars = 95000
	const costPerSingleCar = 10000

	groupsOfTen := carsCount / 10
	remainingCars := carsCount % 10

	totalCost := (groupsOfTen * costPerTenCars) + (remainingCars * costPerSingleCar)
	return uint(totalCost)
}
