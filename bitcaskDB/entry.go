package bitcaskdb

import (
	"bytes"
	"encoding/binary"
	"hash/crc32"
	"time"
)

const (
	CRC32_SIZE      = 4
	TIMESTAMP_SIZE  = 8
	KEY_SIZE        = 4
	VALUE_SIZE      = 4
	KEY_MAX_CHARS   = 2147483647
	VALUE_MAX_CHARS = 2147483647
)

type Entry struct {
	crc32     uint32
	timestamp int64
	keySize   int
	valueSize int
	key       string
	value     string
	data      []byte
}

func NewEntry(key, value string) *Entry {
	timestamp := time.Now().Unix()
	keySize := len(key)
	valueSize := len(value)

	data := toBytes(timestamp, keySize, valueSize, key, value)
	crc32Value := crc32.ChecksumIEEE(data)
	return &Entry{
		crc32:     crc32Value,
		timestamp: timestamp,
		keySize:   len(key),
		valueSize: len(value),
		key:       key,
		value:     value,
		data:      data,
	}
}

func toBytes(timestamp int64, keySize, valueSize int, key, value string) []byte {
	buf := new(bytes.Buffer)

	binary.Write(buf, binary.LittleEndian, timestamp)
	binary.Write(buf, binary.LittleEndian, int32(keySize))
	binary.Write(buf, binary.LittleEndian, int32(valueSize))
	buf.Write([]byte(key))
	buf.Write([]byte(value))

	return buf.Bytes()
}

func (e *Entry) Size() int {
	return CRC32_SIZE + TIMESTAMP_SIZE + KEY_SIZE + VALUE_SIZE + e.keySize + e.valueSize
}
