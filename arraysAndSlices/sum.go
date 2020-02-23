package main

//Sum will add all numbers in an array together and return the result
func Sum(numbers []int) int {
	var sum int = 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

//SumAll will take any number of arrays and return the total for each
func SumAll(numbersToSum ...[]int) []int {
	var sums []int //slice of an array with ints
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers)) // sum a slice and then append it to the slice of slices
	}
	return sums

}

//SumAllTails will take any number of arrays and return the total for the tail of each array
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			//add a zero to the sums array if it is empty
			sums = append(sums, 0)
		} else {
			tail := numbers[1:] //slice the numbers arr starting at index 1 to the end of the array
			sums = append(sums, Sum(tail))
		}
	}
	return sums

}

func main() {
}
