package main

import (
	"io"
	"log/slog"
	"net/http"
	"os"

	"go_mdb/db_server"
)

type Handlers struct {
	logger *slog.Logger
}

func (h *Handlers) initLog() {
	h.logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
}

func (h *Handlers) addMessage(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		errorMessage(w, err)
	}

	db_server.InsertMessage(body)

	h.logger.Debug(string(body))
}

func errorMessage(w http.ResponseWriter, err error) {
	w.Write([]byte(err.Error()))
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Header().Set("Content-Type", "application/json")
}

func main() {
	hs := Handlers{}
	hs.initLog()

	http.HandleFunc("/add_message", hs.addMessage)

	http.ListenAndServe(":3000", nil)
}
