/*
// Sample code to perform I/O:

fmt.Scanf("%s", &myname)            // Reading input from STDIN
fmt.Println("Hello", myname)        // Writing output to STDOUT

// Warning: Printing unwanted or ill-formatted data to output will cause the test cases to fail
*/

// Write your code here

package main

import "fmt"

var isPrime [1000001]bool
var commutate [1000001]int

func sieve() {
	for i := 0; i < 1000001; i++ {
		isPrime[i] = true
	}
	isPrime[0] = false
	isPrime[1] = false
	for i := 2; i < 1000001; i++ {
		if isPrime[i] == true {
			for j := i * i; j < 1000001; j += i {
				isPrime[j] = false
			}
		}
	}
	var count int
	count = 0
	for i := 0; i < 1000001; i++ {
		if isPrime[i] == true {
			count++
		}
		if isPrime[count] == true {
			commutate[i] = 1
		}
		if isPrime[count] == false {
			commutate[i] = 0
		}

	}
	for i := 1; i < 1000001; i++ {
		commutate[i] += commutate[i-1]
	}
}

func main() {
	sieve()
	var T int
	fmt.Scanf("%d", &T)
	for T > 0 {
		var L int
		var R int
		fmt.Scanf("%d%d", &L, &R)

		fmt.Println(commutate[R] - commutate[L-1])
		T--
	}
}
