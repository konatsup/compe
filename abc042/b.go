package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, l int
	fmt.Scan(&n, &l)
	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	result := ""
	for _, v := range s {
		result += v
	}
	fmt.Println(result)
}
