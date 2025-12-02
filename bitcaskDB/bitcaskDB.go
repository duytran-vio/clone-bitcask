package bitcaskdb

import (
	"strings"
)

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
	// Placeholder implementation
	return "value_of_" + key, nil
}

func (db *BitcaskDB) Put(key, value string) error {
	// Placeholder implementation
	return nil
}

func (db *BitcaskDB) Delete(key string) error {
	// Placeholder implementation
	return nil
}
