package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/mcwumbly/bowl-kata-pp-01/game"
	"github.com/mcwumbly/bowl-kata-pp-01/request"
	"github.com/mcwumbly/bowl-kata-pp-01/view"
)

func main() {
	var bowls []int
	app := game.Game{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var err error
			bowls, err = app.AddBowl(bowls, bowlFromRequest(r))
			if err != nil {
				panic(err)
			}
		}
		w.Write(response(bowls, app))
	})
	s := http.Server{Addr: ":8080"}
	go func() {
		s.ListenAndServe()
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT)

	fmt.Println("let's start bowling!")
	<-signals
	fmt.Println("shutting down...")
	s.Shutdown(context.Background())
	fmt.Println("exiting...")
}

func bowlFromRequest(r *http.Request) int {
	var bowl request.BowlRequest
	defer r.Body.Close()
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bodyBytes, &bowl)
	if err != nil {
		panic(err)
	}
	return bowl.Bowl.Pins
}

func response(bowls []int, app game.Game) []byte {
	var gameView view.Game
	gameView.Frames = app.Frames(bowls)
	total := 0
	for _, bowl := range bowls {
		total += bowl
	}
	gameView.Total = total
	respBytes, err := json.Marshal(view.Response{Game: gameView})
	if err != nil {
		panic(err)
	}
	return respBytes
}
