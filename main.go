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
	app := game.New()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			bowl, err := bowlFromRequest(r)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err)))
				return
			}
			err = app.AddBowl(bowl)
			if err != nil {
				w.WriteHeader(http.StatusConflict)
				w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err)))
				return
			}
		}
		w.Write(response(app))
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

func bowlFromRequest(r *http.Request) (int, error) {
	var bowl request.BowlRequest
	defer r.Body.Close()
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return -1, err
	}
	err = json.Unmarshal(bodyBytes, &bowl)
	if err != nil {
		return -1, err
	}
	return bowl.Bowl.Pins, nil
}

func response(app *game.Game) []byte {
	var gameView view.Game
	gameView.Frames = app.Frames()
	total := 0
	for _, bowl := range app.Bowls() {
		total += bowl
	}
	gameView.Total = total
	gameView.CurrentFrame = app.CurrentFrame()
	gameView.RemainingPins = app.RemainingPins()
	respBytes, err := json.Marshal(view.Response{Game: gameView})
	if err != nil {
		panic(err)
	}
	return respBytes
}
