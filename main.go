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

	"github.com/mcwumbly/bowl-kata-pp-01/request"
	"github.com/mcwumbly/bowl-kata-pp-01/view"
)

var (
	bowls []int
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			bowls = append(bowls, bowlFromRequest(r))
		}
		w.Write(response(bowls))
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

func response(bowls []int) []byte {
	var gameView view.Game
	if gameView.Frames == nil {
		gameView.Frames = []view.Frame{}
	}
	if len(bowls) > 0 {
		total := 0
		balls := []view.Ball{}
		for i, bowl := range bowls {
			balls = append(balls, view.Ball{
				Ball: i + 1,
				Pins: bowl,
			})
			total += bowl
		}
		gameView = view.Game{
			Frames: []view.Frame{{
				Frame: 1,
				Balls: balls,
				Total: total,
			}},
			Total: total,
		}
	}
	respBytes, err := json.Marshal(view.Response{Game: gameView})
	if err != nil {
		panic(err)
	}
	return respBytes
}
