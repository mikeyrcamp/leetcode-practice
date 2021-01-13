package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	triangle := pascalTriangle(10)
	var b strings.Builder
	for _, v := range triangle {
		row := make([]string, len(v))
		for j, w := range v {
			row[j] = strconv.Itoa(w)
		}
		b.WriteString(strings.Join(row, " "))
		b.WriteByte('\n')
	}
	fmt.Println(b.String())
}

func pascalTriangle(levels int) [][]int {
	var triangle [][]int

	for i := 1; i <= levels; i++ {
		triangle = append(triangle, make([]int, i))

		triangle[i-1][0] = 1

		if i > 2 {
			for j := 1; j < i-1; j++ {
				triangle[i-1][j] = triangle[i-2][j-1] + triangle[i-2][j]
			}
		}

		if i >= 2 {
			triangle[i-1][i-1] = 1
		}
	}

	return triangle
}
