package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	d := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Scan(&d[i])
	}

	numSet := map[int]struct{}{}
	for _, v := range d {
		numSet[v] = struct{}{}
	}

	var min int
	for i := 1; i < 10; i++ {
		if _, ok := numSet[i]; !ok {
			min = i
		}
	}

	str := fmt.Sprint(n)
	runes := []rune(str)
	minStr := strconv.Itoa(min)
	minRunes := []rune(minStr)
	if runes[0] >= minRunes[0] {
		fmt.Println(n)
		return
	}
	runes[0] = minRunes[0]
	fmt.Println(string(runes))

}
