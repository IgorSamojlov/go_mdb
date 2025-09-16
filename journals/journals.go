package journals

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
)

const (
	FILEPATH      = "tmp/journal0"
	FILEPATH1     = "tmp/journal1"
	FILEPATH2     = "tmp/journal2"
	SEP           = `|`
	MAX_FILE_SIZE = 100
)

type Journal struct {
	mu          sync.Mutex
	currentFile *os.File
}

func (j *Journal) WriteData(data []byte) error {
	var b bytes.Buffer

	getAt, err := time.Now().MarshalBinary()
	if err != nil {
		return err
	}

	uuid, err := uuid.New().MarshalBinary()
	if err != nil {
		return err
	}

	b.Write(getAt)
	b.Write([]byte(SEP))
	b.Write(uuid)
	b.Write([]byte(SEP))
	b.Write(data)
	b.Write([]byte(`\n`))

	j.mu.Lock()
	defer j.mu.Unlock()

	err = j.writeToFile(b)
	if err != nil {
		return err
	}

	fi, err := j.currentFile.Stat()
	if err != nil {
		return err
	}

	if fi.Size() >= MAX_FILE_SIZE {
		j.switchFile()
	}

	return nil
}

func (j *Journal) writeToFile(data bytes.Buffer) error {
	_, err := j.currentFile.Write(data.Bytes())
	if err != nil {
		return err
	}

	err = j.currentFile.Sync()
	if err != nil {
		return err
	}

	return nil
}

func (j *Journal) switchFile() error {
	fn, err := nextBlankFileName()
	if err != nil {
		return err
	}

	f, err := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}

	j.currentFile = f

	fmt.Println("Switch")

	return nil
}

func nextBlankFileName() (string, error) {
	var journalFile = []string{FILEPATH, FILEPATH1, FILEPATH2}

	for _, fn := range journalFile {
		fi, err := os.Stat(fn)
		if err != nil {
			return "", err
		}

		if fi.Size() == 0 {
			return fn, nil
		}
	}

	return "", nil
}

func New() (*Journal, error) {
	fileName, err := nextBlankFileName()
	if err != nil {
		return nil, err
	}

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return nil, err
	}

	return &Journal{currentFile: f}, nil
}

func (j *Journal) Close() error {
	return j.currentFile.Close()
}
