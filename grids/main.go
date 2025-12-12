// Package grid main motivation is to study chess
package grid

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

type TwoDGrid[T any] struct {
	row, col int
	grid     [][]T
}

func (g *TwoDGrid[T]) Init(row, col int) {
	grid := make([][]T, row)
	g.row = row
	g.col = col
	g.grid = grid

}

func (g *TwoDGrid[T]) ColumnSize() int {
	return len(g.grid[0]) - 1
}

func (g *TwoDGrid[T]) Size() int {
	return len(g.grid)
}

func (g *TwoDGrid[T]) Construct(row, col int, fill ...T) {
	grid := make([][]T, row)
	for i := range row {
		grid[i] = make([]T, col)
		if len(fill) > 0 {
			for j := range col {
				grid[i][j] = fill[j%len(fill)] // repeate column
			}
		}
	}
	g.row = row
	g.col = col
	g.grid = grid
}

func (g *TwoDGrid[T]) Range() [][]T {
	return g.grid
}

// Diag returns nth diagonals like 1,2,3,...
func (g *TwoDGrid[T]) Diag() [][]T {
	d := make([][]T, g.row)
	// root := g.grid
	row, col := g.row, g.col
	for i := range row {
		d[i] = make([]T, col)

	}
	return d
}

// TDDiagonals returns diagonal
func (g *TwoDGrid[T]) TDDiagonals() [][]T {
	diag := make([][]T, g.Size())
	for r := range g.Size() {
		diag[r] = []T{g.grid[r][r]}
	}
	return diag
}

// Diagonals nth diagosnals
func (g *TwoDGrid[T]) Diagonals() []T {
	diag := make([]T, 0, g.ColumnSize())
	for r := range g.ColumnSize() {
		diag = append(diag, g.grid[r][r])
	}
	return diag
}

// DDDiagonals returns the the nth diagonals
// algorithm: run from the len of the root
// decrease the row current track-col+1 same as for loop i<len(array)
// c=col-1 // as we cannot deduce the i we take c=col-1 rather than i=col-1
// [1],[1,2].....
func (g *TwoDGrid[T]) DDDiagonals() [][]T {
	var diag [][]T
	root := g.grid
	row, col := len(root), len(root[0])
	for i := range root {
		var r = 0
		var di []T

		if i >= col {
			r = i - col + 1
		}

		// becuase we cant able to deduce i so done shadowing
		c := i

		if c >= col {
			c = col - 1
		}

		for r < row && c >= 0 {
			di = append(di, root[c][r])
			c--
			r++
		}

		diag = append(diag, di)
	}

	return diag
}

// MatAdd matrix addition
func MatAdd(A, B []int) []int {
	var m []int
	for inx, a := range A {
		m = append(m, B[inx]+a)
	}
	return m
}

// SubAdd matrix subtraction
func SubAdd(A, B []int) []int {
	var m []int
	for inx, a := range A {
		m = append(m, B[inx]+a)
	}
	return m
}

// MatMul AB if you pass rowB than BA
// for a[1,2,3]=> 3x1 and b[1,2,3]=> 3x1
func MatMul(a, b []int, rowA, colA, colB int) []int {
	m := make([]int, rowA*colB)
	for i := range rowA {
		for j := range colB {
			sum := 0
			for k := range colA {
				sum += a[i*colA+k] * b[k*colB+j]
			}
			m[i*colB+j] = sum
		}
	}

	return m
}

// DMatMul two d multiplcation
func DMatMul(a, b [][]int) [][]int {
	rowa, rowb := len(a), len(b)
	cola, colb := len(a[0]), len(b[0])
	res := make([][]int, rowa)
	if cola != rowb {
		return nil
	}
	for i := range res {
		res[i] = make([]int, colb)
		for j := range colb {
			for k := range cola {
				res[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return res
}

// Run a temp file for running in the main func
func Run() {
	board := TwoDGrid[string]{}
	row, col := 4, 4
	board.Construct(row, col, ".")

	cols := make(map[int]bool)
	d1 := make(map[int]bool)
	d2 := make(map[int]bool)
	var s [][]string
	CalcNQueen(board, 0, "Q", ".", cols, d1, d2, &s)
	fmt.Println(s, board.Size())
	// lowerBound := sort.Search(len(x), func(i int) bool {
	// 	return x[i] >= 50
	// })
	// upperBound := sort.Search(len(x), func(i int) bool {
	// 	return x[i] > 50
	// })
	data := []string{"a", "b"}

	fmt.Println(data)
}

// Less ideal help for sorting to less
// i basically copied the cpp source code next_prev
func Less(data sort.Interface, i, j int) bool {
	return data.Less(i, j)
}

// NextPerm tried implmenting the next prew but its failed we'll do it later
func NextPerm(data []int) bool {
	_Reverse(sort.IntSlice(data), 1, 2)
	slices.Backward(data)
	x := data
	slices.Sort(data)
	y := data
	_s := IsSortedUntil(data)
	if !slices.Equal(x, y) {

		for r := range _s {
			x[r], x[r] = x[r], x[r]
		}
	}
	return !slices.Equal(x, y)
}

// IsSortedUntil name says all
func IsSortedUntil(data []int) int {
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			return i
		}
	}
	return len(data)
}

// _Reverse this is bsacill the cpp next_prev codes the make_revverse_iterator shit
func _Reverse(data sort.Interface, start, end int) {
	for start < end {
		data.Swap(start, end)
		start++
		end--
	}
}

// CalcNQueen [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
// i have used the enumeration technique of backtracking
func CalcNQueen(board TwoDGrid[string], row int, place, empty string,
	cols, diag1, diag2 map[int]bool, sols *[][]string) bool {
	if row == board.Size() {
		s := make([]string, board.Size())
		for r := 0; r < board.Size(); r++ {
			s[r] = strings.Join(board.Range()[r], "")
		}
		*sols = append(*sols, s)

	}

	for col := 0; col < board.Size(); col++ {

		d1 := row + col                      // front
		d2 := row - col + (board.Size() - 1) // reverse
		// if there is no pos to place the queen
		// skip
		if cols[col] || diag1[d1] || diag2[d2] {
			continue
		}

		// place the queen
		board.Range()[row][col] = place
		cols[col] = true
		diag1[d1] = true
		diag2[d2] = true

		if CalcNQueen(board, row+1, empty, place, cols, diag1, diag2, sols) {
			return true
		}

		// backtrack
		board.Range()[row][col] = empty
		cols[col] = false
		diag1[d1] = false
		diag2[d2] = false

	}
	return false
}

// Reverse revers the the number in the array
// i mean this is horrible
func Reverse(_range int) []int {
	var d []int
	for r := _range; r > 0; r++ {
		d = append(d, r)
	}
	return d
}
