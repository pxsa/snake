package main

import (
	"fmt"
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
	value string
}

// snake
type Snake struct {
	length int
	headX int
	headY int
	body []*part
}

// food
type Food struct {
	x,y int
}

// policy
// 4. Moving the snake
// 5. Specifying game over situation

func main() {
	var input []byte = make([]byte, 1)

	// 1. Generating the Field
	field := New(Width, Height)
	field.GenerateWorld()
	
	// 2. Initialize a snake on the Field
	snake := field.InitSnake()
	
	// 3. Generate a random food inside Field
	field.CreateFood()
	
	// core engin of the program
	for string(input) != "q" {
		
		field.Draw()
		os.Stdin.Read(input)
		snake.Move(string(input), field)
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
		f.world[col][row] = "O"
	case "snake":
		f.world[col][row] = "+"
	case "snake-head":
		f.world[col][row] = "*"
	default:
		f.world[col][row] = " "
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

	f.UpdateWorld(x, y, "food")
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
		body: []*part{},
	}
	snake.body = append(snake.body, &part{x, y, "*"})
	field.UpdateWorld(x, y, "snake-head")

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
	if i == 0 || i == 9 || j == 0 || j == 9 {
		// invalid i&j => lose
		return false
	}

	// Move to the next place
	s.headX = i
	s.headY = j
	temp := s.body[0]
	field.UpdateWorld(temp.x, temp.y, "default")

	s.body[0].x = s.headX
	s.body[0].y = s.headY
	field.UpdateWorld(s.headX, s.headY, "snake-head")
	
	return true
}
