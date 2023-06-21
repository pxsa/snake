package main

import (
	"fmt"
)

// CONST
const (
	Width = 10
	Height = 10
)

// field
type Field struct {
	width int
	height int
	world [][]string
}

// snake
type Snake struct {
	length int
	headX int
	headY int
	body []string
}

// food
type Food struct {
	x,y int
}

// policy
// 2. Generate a random food inside Field
// 3. Initialize a snake on the Field
// 4. Moving the snake
// 5. Specifying game over situation

func main() {
	// 1. Generating the Field
	field := New(Width, Height)
	field.GenerateBorder()
	field.Draw()

}



// ----------------------------------------------------------------------------
// Helper functions
func New(width, height int) *Field {
	var field Field
	
	world := make([][]string, height)
	for index := range world {
		world[index] = make([]string, width)
	}

	field.height = height
	field.width = width
	field.world = world 

	return &field
}

func (f *Field) GenerateBorder() {
	for i := 0; i < f.height; i++ {
		for j := 0; j < f.width; j++ {
			if i == 0 || i == f.height-1 || j == 0 || j == f.width-1 {
				f.world[i][j] = "#"
			} else {
				f.world[i][j] = " "
			}
		}
	}
}

func (f *Field) Draw() {
	for _, row := range f.world {
		// if i == 0 || i == f.height-1 {
		// 	s := strings.Join(row, "")
		// 	fmt.Printf("%s\n", s)
		// }
		fmt.Println(row)
	}
}
