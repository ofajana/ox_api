package oxgame

import (
	_ "fmt"
	"github.com/ofajana/ox_api/utils"
	"strings"
)

// Represents a set of coordinates for slice cell
type box struct {
	row int
	col int
}

var boardMap map[int]box

func init() {
	// Creating Mapping of box Id's to actual slice cell coordinates
	boardMap = make(map[int]box)
	boardMap[1] = box{0, 0}
	boardMap[2] = box{0, 1}
	boardMap[3] = box{0, 2}
	boardMap[4] = box{1, 0}
	boardMap[5] = box{1, 1}
	boardMap[6] = box{1, 2}
	boardMap[7] = box{2, 0}
	boardMap[8] = box{2, 1}
	boardMap[9] = box{2, 2}
}

// New Returns Pointer To a New Game
func New() *Game {
	randString := utils.RandomString(5)
	squares := make([][]boardSquare, 3)
	moves := make(map[string]int)
	for idx := 0; idx < 3; idx++ {
		squares[idx] = getBoardSlice(3)
	}

	return &Game{
		Identifier:   randString,
		nextPlayer:   "X",
		gameSquares:  squares,
		moves:        moves,
		blankSquares: 9,
	}

}

// Initialise Game Board row
func getBoardSlice(size int) []boardSquare {
	val := make([]boardSquare, size)
	for idx := 0; idx < size; idx++ {
		val[idx] = boardSquare{}
	}

	return val
}

// Play function to register a move on the game board
func Play(player string, game *Game, boxNo int) string {
	_box, ok := boardMap[boxNo]
	if !ok {
		return "Invalid Board Square Provided"
	}

	return game.play(strings.ToUpper(player), _box.row, _box.col)
}

// Result returns outcome/result of  a Game
func Result(game *Game) string {

	return game.getResult()
}

// State returns current state of a Game
func State(game *Game) string {

	return game.getState()
}
