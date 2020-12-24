package main

import "fmt"

func main() {
	var slc =  [][]int{
		{1,2,3,4},
		{12,13,14,5},
		{11,16,15,6},
		{10,9,8,7},
	}
	if isValid(slc) {
		fmt.Printf("%d",matrix(slc))
	}else {
		fmt.Print("Bad demension, or empty matrix")
	}
}

func isValid(matrix [][]int) bool{
	if len(matrix) == 0{
		return false
	}
	for _, f := range matrix {
		if len(f) != len(matrix) {
			return false
		}
	}
	return true

}

func matrix(matrix [][]int)  []int {
	var result []int
	size := len(matrix)
	start := 0
	end := size - 1
	for end>0 {
		for c := start; c<=end;c++ {
			result = append(result,matrix[start][c])
		}
		for c := start+1; c<=end;c++ {
			result = append(result,matrix[c][end])
		}
		for c := end - 1; c >= start;c-- {
			result = append(result,matrix[end][c])
		}
		for c := end-1; c>=start+1;c-- {
			result = append(result,matrix[c][start])
		}
		end --
		start ++
	}

	return result
}
