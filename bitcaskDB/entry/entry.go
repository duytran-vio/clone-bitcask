package entry

import (
	"encoding/binary"
	"hash/crc32"
	"time"
)

const (
	CRC32_SIZE       = 4
	TIMESTAMP_SIZE   = 8
	KEY_FIELD_SIZE   = 4
	VALUE_FIELD_SIZE = 4
	KEY_MAX_CHARS    = 1024
	VALUE_MAX_CHARS  = 1 << 20
)

type Entry struct {
	timestamp int64
	keySize   int
	valueSize int
	key       string
	value     string
}

func NewEntry(key, value string) *Entry {
	if len(key) > KEY_MAX_CHARS {
		panic("key size exceeds maximum limit")
	}
	if len(value) > VALUE_MAX_CHARS {
		panic("value size exceeds maximum limit")
	}

	if key == "" || value == "" {
		panic("key and value must be non-empty")
	}

	return &Entry{
		timestamp: time.Now().Unix(),
		keySize:   len(key),
		valueSize: len(value),
		key:       key,
		value:     value,
	}
}

func (e *Entry) Encode() []byte {
	data := make([]byte, e.Size())

	// Layout:
	// [0:4]   CRC32
	// [4:12]  timestamp (int64)
	// [12:16] key size (uint32)
	// [16:20] value size (uint32)
	// [20:...] key + value

	offset := CRC32_SIZE

	binary.LittleEndian.PutUint64(data[offset:], uint64(e.timestamp))
	offset += TIMESTAMP_SIZE

	binary.LittleEndian.PutUint32(data[offset:], uint32(e.keySize))
	offset += KEY_FIELD_SIZE

	binary.LittleEndian.PutUint32(data[offset:], uint32(e.valueSize))
	offset += VALUE_FIELD_SIZE

	copy(data[offset:], []byte(e.key))
	offset += e.keySize

	copy(data[offset:], []byte(e.value))

	crc32Value := crc32.ChecksumIEEE(data[CRC32_SIZE:])
	binary.LittleEndian.PutUint32(data[0:CRC32_SIZE], crc32Value)

	return data
}

func (e *Entry) Size() int {
	return CRC32_SIZE + TIMESTAMP_SIZE + KEY_FIELD_SIZE + VALUE_FIELD_SIZE + e.keySize + e.valueSize
}
