package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	for range make([]int, k) {
		if n%200 == 0 {
			n /= 200
		} else {
			n = n*1000 + 200
		}
	}
	fmt.Println(n)
}
