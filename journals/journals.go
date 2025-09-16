package journals

import (
	"bytes"
	"encoding/json"
	"os"
	"time"

	"go_mdb/types"

	"github.com/google/uuid"
)

const FILE_PATH = "../tmp/journal0"

func WriteToJournal(message types.Message) error {
	f, err := os.OpenFile(FILE_PATH, os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()

	journalItem := types.JournalItem{
		Message: message,
		GetAt:   time.Now(),
		UUID:    uuid.New(),
	}

	bData, err := json.Marshal(journalItem)
	if err != nil {
		return err
	}

	bytes := bytes.NewBuffer(bData)
	bytes.WriteString("\n")

	_, err = f.Write(bData)
	if err != nil {
		return err
	}

	err = f.Sync()
	if err != nil {
		return err
	}

	f.Close()

	return nil
}
