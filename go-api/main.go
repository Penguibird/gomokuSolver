package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

const ColCount = 4;
const RowCount = 4;
var g = Game{
	currentState: nil,
	startingPlayer: Circles,
}

func main() {

	http.HandleFunc("/newGame", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		g = Game{
			currentState: nil,
			startingPlayer: Circles,
		}
		w.Write([]byte("SUCCESS"))
	})

	http.HandleFunc("/turn", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		var t struct {
			Row    int
			Column int
		}
		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			fmt.Println(err)
			return
		}
		g.MakeOpponentMove(t.Row, t.Column)
		// g.startingPlayer = g.startingPlayer.Toggle()

		// fmt.Println(g)
		g.MakeNextMove()

		g.ServeFieldList(w)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
	})
	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))

}
