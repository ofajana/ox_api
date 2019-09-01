package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ofajana/ox_api/oxgame"
	"log"
	"net/http"
	"strconv"
)

var port string
var games map[string]*oxgame.Game

func init() {
	games = make(map[string]*oxgame.Game)
	flag.StringVar(&port, "p", "8080", "Specify Port For Api Server")
	flag.Parse()
	port = fmt.Sprintf(":%v", port)

}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", startGame)
	router.HandleFunc("/games/result/{id}", result).Methods("GET")
	router.HandleFunc("/games/state/{id}", state).Methods("GET")
	router.HandleFunc("/play/{id}/{player}/{box}", play).Methods("GET")

	log.Fatal(http.ListenAndServe(port, router))

}

func result(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	game, ok := games[id]
	if !ok {
		respMsg := oxgame.Message{
			GameId:  id,
			Topic:   "Game Result",
			Message: "No Game With Id Provided"}
		json.NewEncoder(w).Encode(respMsg)
		return
	}

	gameResult := oxgame.Result(game)
	respMsg := oxgame.Message{
		GameId:  id,
		Topic:   "Game Result",
		Message: gameResult}
	json.NewEncoder(w).Encode(respMsg)

}

func play(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	pl := mux.Vars(r)["player"]
	box := mux.Vars(r)["box"]

	game, ok := games[id]

	if !ok {
		respMsg := oxgame.Message{
			GameId:  id,
			Topic:   "Play Turn",
			Message: "No Game with Provided Game Id",
		}
		json.NewEncoder(w).Encode(respMsg)
		return
	}

	if _, err := strconv.Atoi(box); err != nil {
		respMsg := oxgame.Message{
			GameId:  id,
			Topic:   "Play Turn",
			Message: "Box Parameter has to be an integer between 1 and 9",
		}
		json.NewEncoder(w).Encode(respMsg)
		return
	}

	boxInt, _ := strconv.Atoi(box)
	playMsg := oxgame.Play(pl, game, boxInt)
	respMsg := oxgame.Message{
		GameId:  id,
		Topic:   "Play Turn",
		Message: playMsg,
	}
	json.NewEncoder(w).Encode(respMsg)

}

func startGame(w http.ResponseWriter, r *http.Request) {
	game := oxgame.New()
	games[game.Identifier] = game
	respMsg := oxgame.Message{
		GameId:  game.Identifier,
		Topic:   "Start Game",
		Message: "New Game Successfuly Initiated"}
	json.NewEncoder(w).Encode(respMsg)

}

func state(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	game, ok := games[id]

	if !ok {
		respMsg := oxgame.Message{
			GameId:  id,
			Topic:   "Game's Current State",
			Message: "No Game With Id Provided"}
		json.NewEncoder(w).Encode(respMsg)
		return
	}
	currentState := oxgame.State(game)
	respMsg := oxgame.Message{
		GameId:  id,
		Topic:   "Game's Current State",
		Message: currentState}
	json.NewEncoder(w).Encode(respMsg)

}

func test() {
	a := oxgame.New()

	fmt.Println(oxgame.Play("X", a, 11))
	fmt.Println(oxgame.Play("X", a, 3))
	fmt.Println(oxgame.Play("X", a, 1))
	fmt.Println(oxgame.Play("O", a, 1))
	fmt.Println(oxgame.Play("O", a, 2))
	fmt.Println(oxgame.Play("x", a, 5))
	fmt.Println(oxgame.Play("O", a, 8))
	fmt.Println(oxgame.Play("x", a, 7))
	fmt.Println(oxgame.Play("O", a, 6))
	fmt.Println(oxgame.Play("x", a, 6))
	fmt.Println(oxgame.Play("x", a, 9))
	fmt.Println(oxgame.Play("o", a, 7))
	fmt.Println(oxgame.Play("o", a, 5))

	fmt.Println(oxgame.Play("x", a, 5))
	fmt.Println(oxgame.Play("x", a, 5))
	fmt.Println("---Testing")

}
