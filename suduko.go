package main

import "fmt"
import "errors"

type Cell struct {
	cell []int
}

type Grid struct {
	size   int
	Cells  [][]Cell
	before [][]int
}

func NewGrid(size int) *Grid {
	c := make([][]Cell, size)
	for j, _ := range c {
		for i := 0; i < size; i++ {
			var instance Cell
			c[j] = append(c[j], instance)
		}
	}
	g := Grid{size, c, nil}
	return &g
}

func (grid *Grid) cellSize() (size int) {
	return 3
}

func (grid *Grid) AddNumber(number, x, y int) (err error) {
	err, found := grid.CheckNumber(number, x, y)
	if err != nil {
		return err
	}
	if !found {
		grid.Cells[x-1][y-1].cell = append(grid.Cells[x-1][y-1].cell, number)
	}
	return err
}

func (grid *Grid) CheckNumber(number, x, y int) (err error, found bool) {
	if x > grid.size || y > grid.size || number > grid.size {
		return errors.New("Out of range"), false
	}
	for _, check := range grid.Cells[x-1][y-1].cell {
		if check == number {
			return nil, true
		}
	}
	return nil, false
}
func (g *Grid) complete(x, y int) bool {

	if len(g.Cells[x-1][y-1].cell) == g.size-1 {
		return true
	}
	return false
}

func (g *Grid) addSmallGrid(number, x, y int) {
	for i := 0; i < 9; i++ {
		cx := x*g.cellSize() + i%g.cellSize()
		cy := y*g.cellSize() + i/g.cellSize()
		if !g.complete(cx+1, cy+1) {
			g.AddNumber(number, cx+1, cy+1)
		}
	}
}

func (g *Grid) CreateGrid(grid [][]int) (err error) {
	if len(grid) > g.size {
		return errors.New("Grid Too large")
	}
	for i, c := range grid {
		if len(c) > g.size {
			return errors.New("Grid Too large")
		}

		for j, cc := range c {

			if cc != 0 {
				g.markCells(cc, i, j)

			}
		}
	}
	for _, c := range grid {
		var a []int
		for _, cc := range c {

			a = append(a, cc)
		}
		g.before = append(g.before, a)

	}
	return nil
}

func (g *Grid) possible(i, j int) (pos []int) {

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, check := range a {
		//find
		found := false
		for _, aa := range g.Cells[i-1][j-1].cell {
			if aa == check {
				found = true
				break
			}
		}
		if !found {
			pos = append(pos, check)
		}
	}
	return pos
}

func (g *Grid) markCells(cc, i, j int) {

	for l := 0; l < g.size; l++ { //Mark row
		if l != j {
			g.AddNumber(cc, i+1, l+1)
		}
	}
	for m := 0; m < g.size; m++ { //Mark Column
		if m != i {
			g.AddNumber(cc, m+1, j+1)
		}
	}

	for k := 0; k < g.size; k++ { //Mark the cell itself
		if cc != k+1 {
			g.AddNumber(k+1, i+1, j+1)
		}
	}
	g.addSmallGrid(cc, i/g.cellSize(), j/g.cellSize()) //Mark the small grid
}

func (grid *Grid) PrintGrid() (s string) {
	for _, c := range grid.Cells {
		s = s + fmt.Sprintf("%v\n", c)
	}
	return s
}

func (g *Grid) missingNumber(c []int) (missing int) {
	for i := 1; i <= g.size; i++ {
		found := false
		for _, n := range c {
			if n == i {
				found = true
			}
		}
		if !found {
			return i
		}
	}
	return 0
}

func (g *Grid) SolveGrid(grid [][]int) (err error, s [][]int) {

	err = g.CreateGrid(grid)
	if err != nil {
		return err, nil
	}
	
	g.findFromMask()
	finished:=false
	

	for i:=0; !finished && i<5;i++  {
	
		g.checkRow()
		g.checkColumn()
		g.checkSmallGrid()
		g.searchForMarksinCell()
		g.singleZero()
		g.findFromMask()
		
		finished=true
		for _, row := range g.before {
			for _,c:= range row{
				if c==0{
					finished =false 
					break
				}
			}
		}
		
	}
	
	
	
	return nil, g.before
}

func (g *Grid) findFromMask() () {

	for i, c := range g.Cells { //Look for cells that could only be one value
		var row []int
		for j, cc := range c {
			if (len(cc.cell) == g.size-1) && !(g.before[i][j] != 0) { //A new value has been found
				number := g.missingNumber(g.Cells[i][j].cell)
				row = append(row, number)
				
				g.before[i][j] = number
				
			} else {
				row = append(row, 0)

			}
		}
		//s = append(s, row)
	}
	//return s
}

func (g *Grid) checkRow() {
	for number := 1; number <= g.size; number++ {
		for i := 0; i < g.size; i++ {
			found := 0
			foundcol := 0
			for j := 0; j < g.size; j++ {
				pos := g.possible(i+1, j+1)
				if contains(pos, number) {
					found++
					foundcol = j
				}
			}

			if found == 1 && (g.before[i][foundcol] == 0) {

				g.markCells(number, i, foundcol)
				g.before[i][foundcol] = number
				//fmt.Println("Found possible i j number",i,foundcol,number,g.before[i][foundcol])
			}
		}
	}
}

func (g *Grid) checkColumn() {
	for number := 1; number <= g.size; number++ {
		for j := 0; j < g.size; j++ {
			found := 0
			foundrow := 0
			for i := 0; i < g.size; i++ {
				pos := g.possible(i+1, j+1)
				if contains(pos, number) {
					found++
					foundrow = i
				}
			}

			if found == 1 && (g.before[foundrow][j] == 0) {

				g.markCells(number, foundrow, j)
				g.before[foundrow][j] = number
				//fmt.Println("Found possible i j number",i,foundcol,number,g.before[i][foundcol])
			}
		}
	}
}

func (g *Grid) checkSmallGrid() {
	for number := 1; number <= g.size; number++ {
		for x := 0; x < g.size; x += g.cellSize() {
			for y := 0; y < g.size; y += g.cellSize() {
				found := 0
				foundx := 0
				foundy := 0
				for i := 0; i < g.size; i++ {
					cx := x/g.cellSize()*g.cellSize() + i%g.cellSize()
					cy := y/g.cellSize()*g.cellSize() + i/g.cellSize()
					pos := g.possible(cx+1, cy+1)
					if contains(pos, number) {
						found++
						foundx = cx
						foundy = cy
					}
				}
				if found == 1 && (g.before[foundx][foundy] == 0) {

					g.markCells(number, foundx, foundy)
					g.before[foundx][foundy] = number
					//fmt.Println("Found possible i j number",i,foundcol,number,g.before[i][foundcol])
				}
			}
		}
	}
}

func (g *Grid) searchForMarksinCell() {
	for number := 1; number <= g.size; number++ {
		for cx := 0; cx < g.size; cx += g.cellSize() {
			for cy := 0; cy < g.size; cy += g.cellSize() {
				dontCheck := false
				for i := 0; i < g.size; i++ {
					if g.before[cx+i/g.cellSize()][cy+i%g.cellSize()] == number {
						dontCheck = true
					}
				}
				if !dontCheck {
					//fmt.Println("Checking",number,cx,cy)
					g.findAndMarkUniqueRow(number, cx, cy)
					g.findAndMarkUniqueColumn(number, cx, cy)

				}
			}
		}
	}
}

func (g *Grid) findAndMarkUniqueRow(number, cx, cy int) {
	found := 0
	foundrow := 0
	for i := 0; i < g.cellSize(); i++ {

		for j := 0; j < g.cellSize(); j++ {
			pos := g.possible(cx+i+1, cy+j+1)
			if contains(pos, number) {

				found++
				foundrow = j
				break
			}

		}
	}
	
	if found == 1 { //Only found in one row
		//fmt.Println(number, "only found in row",foundrow, "of cell", cx,cy)
		
		for l := 0; l < g.size; l++ { //Mark row
			if l/g.cellSize()*g.cellSize() !=cy{
				g.AddNumber(number,cx+foundrow+1,l+1)
				//fmt.Println("Adding",number,foundrow+cx,l,"to mask")
			}
		}
	}
}

func (g *Grid) findAndMarkUniqueColumn(number, cx, cy int) {
	found := 0
	foundcol := 0
	for i := 0; i < g.cellSize(); i++ {

		for j := 0; j < g.cellSize(); j++ {
			pos := g.possible(cx+j+1, cy+i+1)
			if contains(pos, number) {

				found++
				foundcol = i
				break
			}

		}
	}
	if found == 1 { //Only found in one col
		//fmt.Println(number, "only found in col",foundcol, "of cell", cx,cy)
		
		for l := 0; l < g.size; l++ { //Mark row
			if l/g.cellSize()*g.cellSize() !=cx{
				g.AddNumber(number, l+1,cy+foundcol+1)
			//	fmt.Println("Adding",number,l,cy+foundcol,"to mask")
			}
		}
	}
}

func (g *Grid) singleZero() {
	for i, c := range g.Cells { //Look for cells that could only be one value
		//countcol:=0
		//countrow:=0
		col := 0
		row := 0
		var colnum []int
		var rownum []int
		for j, _ := range c {
			if g.before[i][j] == 0 { //An empty cell has been found
				//countrow++
				col = j
			} else {
				colnum = append(colnum, g.before[i][j])
			}
			if g.before[j][i] == 0 { //An empty cell has been found
				//countcol++
				row = j
			} else {
				rownum = append(rownum, g.before[j][i])
			}

		}
		if len(colnum) == 8 { //If only onein the row
			number := g.missingNumber(colnum)
			//fmt.Println("Found col", number, i, col, colnum)
			g.markCells(number, i, col)
			g.before[i][col] = number
		}
		if len(rownum) == 8 { //If only onein the row
			number := g.missingNumber(rownum)
			g.markCells(number, row, i)
			g.before[row][i] = number
			//fmt.Println("Found row", number, row, i, rownum)
		}
	}

}

func contains(list []int, number int) (found bool) {
	for _, n := range list {
		if n == number {
			return true
		}
	}
	return false
}

func main() {
		grid:=[][]int{
			[]int{0,0,8,0,0,4,0,0,9},
			[]int{0,0,1,0,0,0,0,0,0},
			[]int{5,3,0,0,0,9,0,7,0},
			[]int{0,0,0,0,0,0,0,0,7},
			[]int{0,0,0,0,5,6,8,0,0},
			[]int{9,0,6,0,4,0,0,1,0},
			[]int{0,0,0,0,6,0,0,9,1},
			[]int{0,0,4,0,0,8,6,0,5},
			[]int{3,0,0,5,0,0,7,4,0},
		}


	g := NewGrid(9)

	
	
	err,s:=g.SolveGrid(grid)
	if err!=nil  {
		panic(1)
	}
	
	
	
	
	for _, row := range s {
		fmt.Println(row)
	}

}
