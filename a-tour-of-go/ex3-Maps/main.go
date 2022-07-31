package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wrds := strings.Fields(s)
	num := len(wrds)
	ans := make(map[string]int)

	for i := 0; i < num; i++ {
		(ans[wrds[i]])++
	}

	return ans
}

func main() {
	wc.Test(WordCount)
}
