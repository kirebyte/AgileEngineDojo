package main

func main() {

}

func Solve(arr []int) int {
	sum := 0
	for i, n := range arr {
		if isPrime(i) {
			sum += n
		}
	}
	return sum
}

func isPrime(number int) bool {
	if number == 0 || number == 1 {
		return false
	}
	for count, i := 1, 2; i < number; i++ {
		if number%i == 0 {
			count++
			if count > 1 {
				return false
			}
		}
	}
	return true
}
