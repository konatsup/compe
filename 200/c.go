package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	b := make([]int, 200)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		b[a[i]%200]++
	}

	cnt := 0
	for i := 0; i < 200; i++ {
		cnt += (b[i]*(b[i]-1))/2
	}
	fmt.Println(cnt)
}
