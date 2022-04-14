package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	setupCloseHandler()
	startServer("8088")
}

func startServer(port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] Request on %s", r.Method, r.URL.Path)

		errs := make(map[string]string, 0)

		var body interface{}
		if r.ContentLength > 0 {
			err := json.NewDecoder(r.Body).Decode(&body)
			if err != nil {
				errs["body_decode"] = err.Error()
				log.Printf("[ERROR] Failed to decode body: %s", err.Error())
			}
		}

		b, _ := json.Marshal(map[string]interface{}{
			"path":    r.URL.Path,
			"method":  r.Method,
			"time":    time.Now().Format(time.RFC3339),
			"headers": r.Header,
			"query":   r.URL.Query(),
			"body":    body,
			"errors":  errs,
		})

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)

		fmt.Println()
	})

	log.Printf("[SERVER] Started on :%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}

func setupCloseHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c,
		syscall.SIGHUP,
		syscall.SIGTERM,
		syscall.SIGUSR1,
		syscall.SIGUSR2,
		syscall.SIGILL,
		syscall.SIGABRT,
		syscall.SIGQUIT,
		os.Interrupt,
	)

	go func() {
		log.Println("[SIGNAL] Waiting to receive...")
		s := <-c
		log.Printf("[SIGNAL] Received: %v", s)
		os.Exit(0)
	}()
}
