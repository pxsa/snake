package main

import (
	"math/rand"
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
// 3. Initialize a snake on the Field
// 4. Moving the snake
// 5. Specifying game over situation

func main() {
	// 1. Generating the Field
	field := New(Width, Height)
	field.GenerateWorld()
	
	// 2. Generate a random food inside Field
	field.CreateFood()
	field.Draw()

	// core engin of the program
	for {
		
	}

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

func (f *Field) GenerateWorld() {
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
		fmt.Println(row)
	}
}

func (f *Field) UpdateWorld(food *Food) {
	f.world[food.x][food.y] = "O"
}

// ----------------------------------------------------------------------------
// Create Food
func (f *Field) CreateFood() {
	food := Food {
		x: rand.Intn(Width-2)+1,
		y: rand.Intn(Height-2)+1,
	}
	f.UpdateWorld(&food)
}