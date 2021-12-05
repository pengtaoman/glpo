package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!", row, col)
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

var dirs = [4]point{
	{-1, 0},
	{0, -1},
	{},
}

func walk(maze [][]int, start, end point) {
	steps := make([][]int, len(maze))

	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	//Q := []point{start}
	//for len(Q) > 0 {
	//	cur := Q[0]
	//	Q := Q[1:]
	//}
}

type point struct {
	i, j int
}

func main() {
	maze := readMaze("me-one/maze/maze.in")
	for _, row := range maze {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}

	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@", len(maze[2]))
	walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
}
