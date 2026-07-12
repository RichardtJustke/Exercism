package lasagna

const OvenTime = 40  

func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven
	panic("RemainingOvenTime not implemented")
}

func PreparationTime(numberOfLayers int) int {
	return numbersOfLayers*2
	panic("PreparationTime not implemented")
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
	panic("ElapsedTime not implemented")
}
