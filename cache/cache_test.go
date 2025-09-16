package cache

import (
	"encoding/json"
	"testing"

	"go_mdb/types"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMCachWrite(t *testing.T) {
	be := assert.New(t)

	var messages [1000]types.Message

	messages[0] = types.Message{UserId: 10, ChatId: 10, Message: "Hello"}

	c := MCache{
		Messages:    messages,
		counter:     1,
		LastElement: 0,
	}

	newMessage := types.Message{UserId: 11, ChatId: 110, Message: "New hello"}

	data, err := json.Marshal(newMessage)
	require.NoError(t, err)

	err = c.Write(data)
	require.NoError(t, err)

	require.NoError(t, err)

	be.Equal(types.Message{UserId: 10, ChatId: 10, Message: "Hello"}, c.Messages[0])
	be.Equal(newMessage, c.Messages[1])
}
