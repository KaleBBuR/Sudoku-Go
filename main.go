package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"text/template"
	"time"

	socketio "github.com/googollee/go-socket.io"
)

func parseBoard(boardStr string) [9][9]int {
	board := []byte(boardStr)
	var b [9][9]int
	json.Unmarshal(board, &b)
	return b
}

func sudokuWebpageHandler(w http.ResponseWriter, r *http.Request) {
	t, parseErr := template.ParseFiles("templates/home.html")
	if parseErr != nil {
		fmt.Printf("\n\nCan not parse home html file %s\n\n", parseErr.Error())
	}
	t.Execute(w, nil)
}

func main() {
	server, serverErr := socketio.NewServer(nil)

	if serverErr != nil {
		log.Fatal(serverErr)
	}

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("Connected: ", s.ID())
		return nil
	})

	server.OnEvent("/", "check", func(s socketio.Conn, boardString string) {
		board := parseBoard(boardString)
		var boardStruct Board
		boardStruct.cells = board
		if boardGood, location, loc1, loc2 := boardStruct.checkStartingBoard(); !boardGood {
			s.Emit("check board", location, loc1, loc2, board)
		} else {
			s.Emit("check board", location, -1, -1, board)
		}
	})

	server.OnEvent("/", "solve", func(s socketio.Conn, boardString string) {
		board := parseBoard(boardString)
		var boardStruct Board
		boardStruct.cells = board
		boardStruct.solveSudoku()
		s.Emit("solved board", boardStruct.cells)
	})

	server.OnEvent("/", "easy", func(s socketio.Conn) {
		var boardStruct Board
		boardStruct.generateBoard("easy")
		s.Emit("easy board", boardStruct.cells)
	})

	server.OnEvent("/", "medium", func(s socketio.Conn) {
		var boardStruct Board
		boardStruct.generateBoard("medium")
		s.Emit("medium board", boardStruct.cells)
	})

	server.OnEvent("/", "hard", func(s socketio.Conn) {
		var boardStruct Board
		boardStruct.generateBoard("hard")
		s.Emit("hard board", boardStruct.cells)
	})

	server.OnEvent("/", "expert", func(s socketio.Conn) {
		var boardStruct Board
		boardStruct.generateBoard("expert")
		s.Emit("expert board", boardStruct.cells)
	})

	server.OnEvent("/", "random", func(s socketio.Conn) {
		difficulties := []string{"easy", "medium", "hard", "expert"}
		rand.Seed(time.Now().Unix())
		difficulty := difficulties[rand.Intn(len(difficulties))]
		var boardStruct Board
		boardStruct.generateBoard(difficulty)
		s.Emit("random board", boardStruct.cells)
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	http.HandleFunc("/", sudokuWebpageHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
