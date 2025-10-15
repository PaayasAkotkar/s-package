// Package grid main motivation is to study chess
package grid

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
