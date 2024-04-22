package main

import (
	"fmt"
	"os"
)

const (
	Size  = 9
	Empty = '.'
)

func main() {
	input := os.Args[1:]

	if len(input) != Size {
		fmt.Println("Error")
		return
	}

	table := make([][]rune, Size)

	for i, char := range input {
		if len(char) != Size {
			fmt.Println("Error")
			return
		}
		table[i] = []rune(char)
	}

	if !Sudoku(table) {
		fmt.Println("Error")
		return
	}

	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if table[i][j] != Empty {
				for k := 0; k < Size; k++ {
					if k != j && table[i][k] == table[i][j] || k != i && table[k][j] == table[i][j] {
						fmt.Println("Error")
						return
					}
				}
			}
		}
	}

	PrintTable(table)
}

func Sudoku(table [][]rune) bool {
	var row, col int
	if !EmptyCell(table, &row, &col) {
		return true
	}

	for num := '1'; num <= '9'; num++ {
		if Check(table, row, col, num) {
			table[row][col] = num
			if Sudoku(table) {
				return true
			}
		}
		table[row][col] = Empty
	}

	return false
}

func EmptyCell(table [][]rune, row, col *int) bool {
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if table[i][j] == Empty {
				*row, *col = i, j
				return true
			}
		}
	}
	return false
}

func Check(table [][]rune, row, col int, num rune) bool {
	return !ChekRowCol(table, row, col, num) && !ChekBox(table, row-row%3, col-col%3, num)
}

func ChekRowCol(table [][]rune, row, col int, num rune) bool {
	for i := 0; i < Size; i++ {
		if table[row][i] == num || table[i][col] == num {
			return true
		}
	}
	return false
}

func ChekBox(table [][]rune, startRow, startCol int, num rune) bool {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if table[row+startRow][col+startCol] == num {
				return true
			}
		}
	}
	return false
}

func PrintTable(table [][]rune) {
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if table[i][j] < '1' || table[i][j] > '9' && table[i][j] != '.' {
				fmt.Println("Error")
				return
			}
		}
	}
	for _, row := range table {
		for i, cell := range row {
			if i == 8 {
				fmt.Printf("%c", cell)
			} else {
				fmt.Printf("%c ", cell)
			}
		}
		fmt.Println()
	}
}
