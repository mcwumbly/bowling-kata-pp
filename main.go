package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/mcwumbly/bowl-kata-pp-01/view"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		respBytes, err := json.Marshal(view.Response{
			Game: view.Game{
				Frames: []view.Frame{},
			},
		})
		if err != nil {
			panic(err)
		}
		w.Write(respBytes)
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
