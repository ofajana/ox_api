package oxgame

import (
	"fmt"
	"strings"
	"sync"
)

type Message struct {
	GameId  string `json:"GameId"`
	Topic   string `json:"Topic"`
	Message string `json:"Message"`
}

type boardSquare struct {
	row  int
	col  int
	mark string
}

// Game represents a noughts and crosses Game.
type Game struct {
	Identifier    string
	moves         map[string]int
	winningPlayer string
	gameSquares   [][]boardSquare
	nextPlayer    string
	gameOver      bool
	result        bool
	blankSquares  int
	sync.Mutex
}

// Returns current state of calling game
func (a *Game) getState() string {
	a.Lock()
	if a.gameOver {
		return a.getResult()
	}

	xplays := a.moves["X"]
	yplays := a.moves["Y"]
	nextPlayer := a.nextPlayer
	a.Unlock()
	return fmt.Sprintf(`Game is still in Play, 
						and there are %v empty 
						boxes on the Game Board
						'X' has had %v turn(s)
						'Y' has had %v turn(s)
						Next to play '%v'`, a.blankSquares, xplays, yplays, nextPlayer)
}

// Returns result of calling game
// if game is still in play, returns a result indicating
// current state of game
func (a *Game) getResult() string {
	a.Lock()
	if a.gameOver {
		if a.result {
			a.Unlock()
			return fmt.Sprintf("Game was Won By '%v'", a.winningPlayer)
		}
		a.Unlock()
		return fmt.Sprintf("Game Ended in a Draw")
	}
	a.Unlock()
	return fmt.Sprintf("Game Is Still In Play, No Result Yet")

}

func (a *Game) play(player string, row int, col int) string {
	// Check Player is Valid
	a.Lock()
	_player := strings.ToUpper(player)
	if _player != "X" && _player != "O" {
		a.Unlock()
		return "Invalid Player, Use 'X' or 'O'"
	}
	// Check if Game is Over
	if a.gameOver {
		a.Unlock()
		return "This Game has Ended, Please Start a New Game"
	}

	// Check if it's calling players turn
	if a.nextPlayer != player {
		a.Unlock()
		return "Not Your Turn"
	}
	// If Board Square is empty
	// Mark it to indicate calling players turn
	//  check if play is a winning move
	// Else Return appropriate message to player
	if a.gameSquares[row][col].mark == "" {
		a.gameSquares[row][col].mark = player
		a.moves[player] = a.moves[player] + 1
		a.blankSquares = a.blankSquares - 1
		if a.isRowWinner(row, player) ||
			a.isColWinner(col, player) ||
			a.isDiagWinner(player) {
			a.toggleNextPlayer()
			a.winningPlayer = player
			a.result = true
			a.Unlock()
			return "Move Successfully Recorded, Winning Move"
		}

		if a.blankSquares < 1 {
			a.gameOver = true
			a.Unlock()
			return "Game Over, Ended in a Draw"
		}
		a.toggleNextPlayer()
		a.Unlock()
		return "Move Successfully Recorded"

	}
	a.Unlock()
	return "Invalid Play, Try another Square"
}

func (a *Game) toggleNextPlayer() {
	if a.nextPlayer == "X" {
		a.nextPlayer = "O"
		return
	}

	a.nextPlayer = "X"
}

// Check if row has a winning combination

func (a *Game) isRowWinner(row int, player string) bool {
	count := 0
	for _, entry := range a.gameSquares[row] {
		if player == entry.mark {
			count++
		}
	}
	if count > 2 {
		a.gameOver = true
	}
	return count > 2
}

// Check if column has a winning combination
func (a *Game) isColWinner(col int, player string) bool {
	count := 0
	for idx := 0; idx < 3; idx++ {
		if player == a.gameSquares[idx][col].mark {
			count++
		}
	}
	if count > 2 {
		a.gameOver = true
	}
	return count > 2
}

// Check if we have a winning diagonal combination
func (a *Game) isDiagWinner(player string) bool {

	if player == a.gameSquares[0][0].mark &&
		player == a.gameSquares[1][1].mark &&
		player == a.gameSquares[2][2].mark {
		a.gameOver = true
		return true
	}

	if player == a.gameSquares[0][2].mark &&
		player == a.gameSquares[1][1].mark &&
		player == a.gameSquares[2][0].mark {
		a.gameOver = true
		return true
	}

	return false

}
