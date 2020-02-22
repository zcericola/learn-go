package main

//Sum will add all numbers in an array together and return the result
func Sum(numbers []int) int {
	var sum int = 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func main() {
}
