package main

import (
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"

	"go_mdb/journals"
)

const FILEPATH = "tmp/journal0"

type Handlers struct {
	logger  *slog.Logger
	journal journals.Journal
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

		return
	}

	err = h.journal.WriteData(body)
	if err != nil {
		errorMessage(w, err)

		return
	}

	h.logger.Debug(string(body))
}

func errorMessage(w http.ResponseWriter, err error) {
	w.Write([]byte(err.Error()))
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Header().Set("Content-Type", "application/json")
}

func main() {
	journal, err := journals.New()
	if err != nil {
		log.Fatal(err)
	}
	defer journal.Close()

	hs := Handlers{journal: *journal}

	hs.initLog()

	http.HandleFunc("/add_message", hs.addMessage)

	http.ListenAndServe(":3000", nil)
}
