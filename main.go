package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const VERSION = "v1.2.4"
const FEATURE_QS_SLEEP = "fs-sleep"
const FEATURE_QS_STATUS = "fs-status"

func main() {
	log.Printf("[SERVER] Fake Server %s", VERSION)
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
				errs["body-decode"] = err.Error()
				log.Printf("[ERROR] Failed to decode body: %s", err.Error())
			}

			defer r.Body.Close()
		}

		// Custom sleep time
		if toSleep := r.URL.Query().Get(FEATURE_QS_SLEEP); toSleep != "" {
			sleepTime, err := time.ParseDuration(toSleep + "ms")

			if err == nil {
				log.Printf("[SLEEP] Want to dream %s", sleepTime)
				time.Sleep(sleepTime)
				log.Printf("[SLEEP] End of the dream %s", sleepTime)
			} else {
				errs["fs-sleep"] = err.Error()
				log.Printf("[ERROR] Failed to parse custom sleep: %s", err.Error())
			}
		}

		// Custom status code
		statusCode := http.StatusOK
		if customStatus := r.URL.Query().Get(FEATURE_QS_STATUS); customStatus != "" {
			parsedStatus, err := strconv.Atoi(customStatus)

			if err == nil {
				statusCode = parsedStatus
				log.Printf("[STATUS] Custom status: %d", parsedStatus)
			} else {
				errs["fs-status"] = err.Error()
				log.Printf("[ERROR] Failed to parse custom status: %s", err.Error())
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
		w.WriteHeader(statusCode)
		_, err := w.Write(b)

		if err != nil {
			log.Printf("[ERROR] Failed write response to client: %s", err.Error())
		}

		fmt.Println()
	})

	log.Printf("[SERVER] Started on :%s\n\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}

func setupCloseHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c,
		syscall.SIGHUP,
		syscall.SIGTERM,
		syscall.SIGILL,
		syscall.SIGABRT,
		syscall.SIGQUIT,
		os.Interrupt,
	)

	go func() {
		log.Printf("[SIGNAL] Waiting to receive...")
		s := <-c
		log.Printf("[SIGNAL] Received: %v", s)
		os.Exit(0)
	}()
}
