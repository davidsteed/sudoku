package main

import "testing"
import "fmt"

func TestSuduko(t *testing.T) {

	g := NewGrid(9)
	if g == nil {
		t.Fatal("Failed to initialise Grid")
	}

	err := g.AddNumber(10, 9, 9)
	if err == nil {
		t.Fatal(err)
	}
	g.AddNumber(1, 9, 9)
	err = g.AddNumber(1, 9, 9)
	t.Log(g.PrintGrid())
}
func TestMissingNumber(t *testing.T) {
	g := NewGrid(9)
	c := Cell{[]int{1, 2, 3, 4, 5, 7, 8, 9}}
	missing := g.missingNumber(c.cell)
	if missing != 6 {
		t.Fatal("Missing Number should be 6 is ", missing)
	}
	t.Log("Missing Number is", missing)

}




func TestCreateGrid(t *testing.T) {

	grid := [][]int{
		[]int{0, 0, 8, 0, 0, 4, 0, 0, 9},
		[]int{0, 0, 1, 0, 0, 0, 0, 0, 0},
		[]int{5, 3, 0, 0, 0, 9, 0, 7, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 7},
		[]int{0, 0, 0, 0, 5, 6, 8, 0, 0},
		[]int{9, 0, 6, 0, 4, 0, 0, 1, 0},
		[]int{0, 0, 0, 0, 6, 0, 0, 9, 1},
		[]int{0, 0, 4, 0, 0, 8, 6, 0, 5},
		[]int{3, 0, 0, 5, 0, 0, 7, 4, 0},
	}

	g := NewGrid(9)

	err := g.CreateGrid(grid)
	if err != nil {
		t.Fatal(err)
	}
	var s string
	for _, c := range grid {
		for _, cc := range c {
			a := (cc != 0)
			s = s + fmt.Sprintf("%v ", a)
		}
		s = s + "\n"
	}
	t.Log(s)
	for _, row := range g.Cells {
		t.Log(row)
	}
}

func TestPossible(t *testing.T) {
	g := NewGrid(9)
	grid := [][]int{
		[]int{0, 0, 8, 0, 0, 4, 0, 0, 9},
		[]int{0, 0, 1, 0, 0, 0, 0, 0, 0},
		[]int{5, 3, 0, 0, 0, 9, 0, 7, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 7},
		[]int{0, 0, 0, 0, 5, 6, 8, 0, 0},
		[]int{9, 0, 6, 0, 4, 0, 0, 1, 0},
		[]int{0, 0, 0, 0, 6, 0, 0, 9, 1},
		[]int{0, 0, 4, 0, 0, 8, 6, 0, 5},
		[]int{3, 0, 0, 5, 0, 0, 7, 4, 0},
	}
	g.SolveGrid(grid)
	all := append(g.Cells[0][0].cell, g.possible(1, 1)...)
	if len(all) != g.size {
		t.Fatal("possible not working")
	}
}

func TestFindAndMarkUniqueRow(t *testing.T){
	grid := [][]int{
		[]int{1, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 1, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 1, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 6, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 2, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 5, 0, 0, 0, 0, 0, 0, 0},
	}

	g := NewGrid(9)

	g.CreateGrid(grid)
	for _, row := range g.before {
		t.Log(row)
	}
	
	g.searchForMarksinCell()
		
	g.findFromMask()
	t.Log("marksincells*****************")
	for _, row := range g.before {
		t.Log(row)
	}
}


func TestSolveGrid(t *testing.T) {

	//	grid:=[][]int{
	//		[]int{0,0,8,0,0,4,0,0,9},
	//		[]int{0,0,1,0,0,0,0,0,0},
	//		[]int{5,3,0,0,0,9,0,7,0},
	//		[]int{0,0,0,0,0,0,0,0,7},
	//		[]int{0,0,0,0,5,6,8,0,0},
	//		[]int{9,0,6,0,4,0,0,1,0},
	//		[]int{0,0,0,0,6,0,0,9,1},
	//		[]int{0,0,4,0,0,8,6,0,5},
	//		[]int{3,0,0,5,0,0,7,4,0},
	//	}
	
	grid := [][]int{
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 8, 0, 0, 0, 3, 0, 0},
		[]int{1, 0, 0, 4, 0, 9, 0, 0, 2},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{8, 0, 7, 0, 0, 0, 6, 0, 9},
		[]int{0, 4, 5, 6, 0, 7, 2, 1, 0},
		[]int{0, 0, 9, 7, 1, 3, 8, 0, 0},
		[]int{0, 0, 0, 0, 9, 0, 0, 0, 0},
		[]int{4, 1, 0, 8, 0, 6, 0, 5, 7},
	}
	done := [][]int{
		[]int{2, 7, 4, 3, 6, 8, 5, 9, 1},
		[]int{5, 9, 8, 2, 7, 1, 3, 4, 6},
		[]int{1, 3, 6, 4, 5, 9, 7, 8, 2},
		[]int{3, 6, 1, 9, 8, 2, 4, 7, 5},
		[]int{8, 2, 7, 1, 4, 5, 6, 3, 9},
		[]int{9, 4, 5, 6, 3, 7, 2, 1, 8},
		[]int{6, 5, 9, 7, 1, 3, 8, 2, 4},
		[]int{7, 8, 2, 5, 9, 4, 1, 6, 3},
		[]int{4, 1, 3, 8, 2, 6, 9, 5, 7},
	}

	g := NewGrid(9)

	err,s:=g.SolveGrid(grid)
	if err!=nil || checkerr(g,done,"solve grid") {
		t.Fatal()
	}
	
	
	
	
	for _, row := range s {
		t.Log(row)
	}
}

func checkerr(g *Grid,done [][]int,desc string) (err bool) {
	for i, c := range g.before {
		for j, cc := range c {
			if cc != 0 && cc != done[i][j] {
				fmt.Println("Cell", i, j, "is", cc, "should be", done[i][j],desc)
				for n,x:=range g.Cells{
					fmt.Print("**",n)
					for z,y:=range x{
						fmt.Print(z,y)
					}
					fmt.Print("\n")
				}
				return true

			}
		}
	}
	return false
}
