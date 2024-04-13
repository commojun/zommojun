package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/commojun/zommojun/config"
)

type API struct {
	Pattern string
	HFunc   func(http.ResponseWriter, *http.Request)
}

func Start() error {
	apis := []API{
		{
			"GET /ping",
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				fmt.Fprintln(w, "OK")
			},
		},
	}

	for _, api := range apis {
		http.HandleFunc(api.Pattern, api.HFunc)
	}

	log.Println("Server Start")
	err := http.ListenAndServe(fmt.Sprintf(":%s", config.ServerPort), nil)
	if err != nil {
		return fmt.Errorf("in server.Start: %w", err)
	}
	return nil
}
