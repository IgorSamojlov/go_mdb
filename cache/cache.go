package cache

import (
	"encoding/json"

	"go_mdb/types"
)

const MAX_SIZE = 1000

type MCache struct {
	Messages    [MAX_SIZE]types.Message
	counter     int
	LastElement int
}

func (mc *MCache) Write(data []byte) error {
	if mc.LastElement == MAX_SIZE-1 {
		mc.LastElement = 0
	}

	mc.LastElement++

	var m types.Message

	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}

	mc.Messages[mc.LastElement] = m

	mc.counter += mc.LastElement + 1

	return nil
}
