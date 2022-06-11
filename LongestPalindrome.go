package main

import "fmt"

func main() {

	var input string
	fmt.Scanf("%s", &input)

	var start int
	var end int
	size := len(input)
	table := make([][]int, size+1)

	for i := 0; i < size; i++ {
		table[i] = make([]int, size)
	}

	for i := 0; i < size; i++ {
		table[i][i] = 1
	}
	for i := 0; i < size-1; i++ {
		if input[i] == input[i+1] {
			table[i][i+1] = 1
			start = i
			end = 2
		}

	}
	for k := 3; k <= size; k++ {
		for i := 0; i < size-k+1; i++ {

			j := i + k - 1
			if table[i+1][j-1] == 1 && input[i] == input[j] {
				table[i][j] = 1
				if k > end {
					start = i
					end = k
				}

			}

		}
	}

	for i := start; i <= end; i++ {
		fmt.Printf("%c", input[i])
	}

}
