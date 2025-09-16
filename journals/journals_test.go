package journals

import (
	"encoding/json"
	"os"
	"testing"

	"go_mdb/types"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecodeShow(t *testing.T) {
	be := assert.New(t)

	message := types.Message{UserId: 10, ChatId: 10, Message: "Hello"}

	err := WriteToJournal(message)
	require.NoError(t, err)

	data, err := os.ReadFile("../tmp/journal0")
	require.NoError(t, err)

	items := []types.JournalItem{}
	json.Unmarshal(data, &items)

	be.Equal(nil, items)
}

func ClearFile(filepath string) error {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}

	defer file.Close()

	return nil
}
