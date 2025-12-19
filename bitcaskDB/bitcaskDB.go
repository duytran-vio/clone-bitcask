package bitcaskdb

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	FILE_BASE_PATH  = "./data/"
	TOMBSTONE_VALUE = "__TOMBSTONE__"
)

var ErrKeyNotFound = errors.New("key not found")

type BitcaskDB struct {
	keyDir            map[string]KeyDirValue // Simplified key directory
	files             []*BitcaskFile
	currentActiveFile *BitcaskFile
}

func NewBitcaskDB() *BitcaskDB {
	newActiveFile := NewBitcaskFile(0, "")
	return &BitcaskDB{
		keyDir:            make(map[string]KeyDirValue),
		files:             []*BitcaskFile{newActiveFile},
		currentActiveFile: newActiveFile,
	}
}

func (db *BitcaskDB) HandleCommand(input string) string {
	parts := strings.Fields(input) // splits on whitespace
	if len(parts) == 0 {
		return ""
	}
	command := parts[0]
	switch command {
	case "GET":
		if len(parts) != 2 {
			return "Usage: GET <key>"
		}
		key := parts[1]
		value, err := db.Get(key)
		if err != nil {
			if errors.Is(err, ErrKeyNotFound) {
				return err.Error() + "\nPlease input a existing key."
			}
			return "Error:" + err.Error()
		} else {
			return "Value:" + value
		}
	case "PUT":
		if len(parts) != 3 {
			return "Usage: PUT <key> <value>"
		}
		key := parts[1]
		value := parts[2]
		err := db.Put(key, value)
		if err != nil {
			return "Error:" + err.Error()
		} else {
			return "OK"
		}
	case "DELETE":
		if len(parts) != 2 {
			return "Usage: DELETE <key>"
		}
		key := parts[1]
		err := db.Delete(key)
		if err != nil {
			return "Error:" + err.Error()
		} else {
			return "OK"
		}
	}
	return "Unknown command"
}

func (db *BitcaskDB) Get(key string) (string, error) {
	keyDirValue, exists := db.keyDir[key]
	if !exists {
		return "", fmt.Errorf("get %q: %w", key, ErrKeyNotFound)
	}

	readFile := db.files[keyDirValue.FileID]
	entryData, err := readFile.ReadAt(keyDirValue.Position, keyDirValue.Size)
	if err != nil {
		return "", err
	}

	return string(entryData), nil
}

func (db *BitcaskDB) Put(key, value string) error {
	if value == TOMBSTONE_VALUE {
		return fmt.Errorf("value cannot be tombstone value")
	}
	newEntry := NewEntry(key, value)
	if (db.currentActiveFile == nil) || (!db.currentActiveFile.canAppend(newEntry.Size())) {
		newFileID := len(db.files)
		newFilePath := FILE_BASE_PATH + "datafile_" + strconv.Itoa(newFileID) + ".db"
		newActiveFile := NewBitcaskFile(newFileID, newFilePath)
		db.files = append(db.files, newActiveFile)
		db.currentActiveFile = newActiveFile
	}

	entryPosition := db.currentActiveFile.Size
	err := db.currentActiveFile.Append(newEntry.encode())
	if err != nil {
		return err
	}

	db.keyDir[key] = KeyDirValue{
		FileID:   db.currentActiveFile.FileID,
		Position: entryPosition,
		Size:     newEntry.Size(),
	}
	return nil
}

func (db *BitcaskDB) Delete(key string) error {
	err := db.Put(key, TOMBSTONE_VALUE)
	if err != nil {
		return err
	}
	delete(db.keyDir, key)
	return nil
}
