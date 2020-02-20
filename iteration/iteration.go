package iteration

//Repeat repeats a string character
func Repeat(character string, numOfTimes int) string {
	if numOfTimes < 1 {
		numOfTimes = 1
	}

	var repeated string
	for i := 0; i < numOfTimes; i++ {
		repeated += character
	}
	return repeated
}
