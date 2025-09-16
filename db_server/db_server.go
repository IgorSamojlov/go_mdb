package db_server

import (
	"encoding/json"

	"go_mdb/journals"
	"go_mdb/types"
)

func InsertMessage(rawMessage []byte) error {
	var message types.Message

	err := json.Unmarshal(rawMessage, &message)
	if err != nil {
		return err
	}

	err = journals.WriteToJournal(message)
	if err != nil {
		return err
	}

	return nil
}
