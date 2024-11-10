package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/frolosofsky/aroundhome/cmd/matcher/api"
	"github.com/frolosofsky/aroundhome/pkg/store/driver/pg"
)

func main() {
	dbconn := os.Getenv("dbconn")
	if len(dbconn) == 0 {
		die("env dbconn is required\n")
	}

	bind := os.Getenv("bind")
	if len(bind) == 0 {
		bind = "127.0.0.1:8080"
	}

	store, err := pg.NewDriver(dbconn)
	if err != nil {
		die("failed to init database: %s\n", err)
	}

	svc := api.Service{
		PartnerStore: store,
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		if err := store.Health(); err != nil {
			log.Printf("[E] Health failed: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	})
	http.HandleFunc("/match", svc.HandleMatch)
	http.HandleFunc("/partners/", svc.HandlePartnerDetails)

	log.Printf("[info] HTTP server starts listening on %s", bind)
	if err := http.ListenAndServe(bind, nil); err != nil {
		die("server failed: %s\n", err)
	}
}

func die(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}
