package main

import "fmt"

var isPrime [86028121]bool
var finalPrime []int32

func sieve() {
	for i := range isPrime {
		isPrime[i] = true
	}
	isPrime[0] = false
	isPrime[1] = false
	for i := 2; i < 86028121; i++ {
		if isPrime[i] == true {
			finalPrime = append(finalPrime, int32(i))
			for j := i * i; j < 86028121; j += i {
				isPrime[j] = false
			}
		}
	}
}

func main() {
	sieve()
	var T int64
	fmt.Scan(&T)
	for T > 0 {
		T--
		var N int64
		fmt.Scan(&N)
		fmt.Println(finalPrime[int32(N-1)])

	}
}
