package birdwatcher

func TotalBirdCount(birdsPerDay []int) int {
	total := 0
	for _, count := range birdsPerDay {
		total += count
	}
	return total
}

func BirdsInWeek(birdsPerDay []int, week int) int {
	start := (week - 1) * 7
	end := start + 7
	if end > len(birdsPerDay) {
		end = len(birdsPerDay)
	}
	return TotalBirdCount(birdsPerDay[start:end])
}

func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
		if i%2 == 0 {
			birdsPerDay[i]++
		}
	}
	return birdsPerDay
}
