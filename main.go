package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
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

// part
type part struct {
	x int
	y int
}

// snake
type Snake struct {
	length int
	headX int
	headY int
	body []*part
}

// policy

func main() {
	var input []byte = make([]byte, 1)
	
	// 1. Generating the Field
	field := New(Width, Height)
	field.GenerateWorld()
	
	// 2. Initialize a snake on the Field
	snake := field.InitSnake()
	
	// 3. Generate a random food inside Field
	field.CreateFood()
	
	// core engine of the program
	for string(input) != "q" {
		
		field.Draw()
		os.Stdin.Read(input)
		// 4. Moving the snake
		// 5. Specifying game over situation
		if !snake.Move(string(input), field) {
			log.Println("Game Over")
			return
		}
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
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

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
	f.ClearWorld()
	for _, row := range f.world {
		fmt.Println(row)
	}
}

func (f *Field) UpdateWorld(row, col int, kind string) {
	switch kind {
	case "food":
		f.world[row][col] = "o"
	case "snake-body":
		f.world[row][col] = "+"
	case "snake-head":
		f.world[row][col] = "x"
	case "default":
		f.world[row][col] = " "
	default:
		f.world[row][col] = " "
	}
}

func (f *Field) ClearWorld() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

// ----------------------------------------------------------------------------
// Create Food
func (f *Field) CreateFood() {
	// check for borders
	x := rand.Intn(Width-2)+1
	y := rand.Intn(Height-2)+1

	// check for snake body
	for f.world[x][y] != " " {
		x = rand.Intn(Width-2)+1
		y = rand.Intn(Height-2)+1
	}

	f.UpdateWorld(y, x, "food")
}

// ----------------------------------------------------------------------------
// Snake
func (field *Field) InitSnake() *Snake{
	x := rand.Intn(Width-2)+1
	y := rand.Intn(Height-2)+1
	snake := &Snake{
		headX: x,
		headY: y,
		length: 1,
	}
	snake.body = append(snake.body, &part{x, y})
	field.UpdateWorld(y, x, "snake-head")

	return snake
}

func (s *Snake) Move(input string, field *Field) bool{
	input = strings.ToLower(input)

	i := s.headX
	j := s.headY

	switch input {
	case "w":
		// Move Up
		j--

	case "a":
		// Move Left
		i--

	case "d":
		// Move Right
		i++

	case "s":
		// Move Down
		j++

	default:
		// Move No-where
	}

	// Validate i & j
	if i == 0 || i == Width-1 || j == 0 || j == Height-1 {
		// invalid i&j => lose
		return false
	}

	// Move to the next place
	s.headX   = i
	s.headY   = j
	lastPart := s.body[s.length-1]
	
	for index := s.length-2; index >= 0; index-- {
		s.body[index+1] = s.body[index]
	}
	
	newPart  := part {x: s.headX, y: s.headY}
	s.body[0] = &newPart

	// Check whether the next place is food or not
	if field.world[s.headY][s.headX] == "o" {
		s.body = append(s.body, lastPart)
		s.length++
		field.CreateFood()
	} else {
		// lastPart = nil
		field.UpdateWorld(lastPart.y, lastPart.x, "default")
	}

	field.UpdateWholdWorld(s)
	return true
}

func(f *Field) UpdateWholdWorld(snake *Snake) {
	var row, col int
	for index, part := range snake.body {
		row = part.y
		col = part.x
		if index == 0 {
			f.UpdateWorld(row, col, "snake-head")
			continue
		}
		f.UpdateWorld(row, col, "snake-body")
	}
}
