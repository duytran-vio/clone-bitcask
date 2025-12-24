package entry

import (
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntry_Encode(t *testing.T) {
	key := "testKey"
	value := "testValue"
	entry := NewEntry(key, value)
	encoded := entry.Encode()

	assert.NotNil(t, encoded)

	expectedSize := CRC32_SIZE + TIMESTAMP_SIZE + KEY_FIELD_SIZE + VALUE_FIELD_SIZE + len(key) + len(value)
	assert.Equal(t, expectedSize, len(encoded))

	offset := CRC32_SIZE + TIMESTAMP_SIZE
	keySize := int(binary.LittleEndian.Uint32(encoded[offset : offset+KEY_FIELD_SIZE]))
	assert.Equal(t, len(key), keySize)

	offset += KEY_FIELD_SIZE
	valueSize := int(binary.LittleEndian.Uint32(encoded[offset : offset+VALUE_FIELD_SIZE]))
	assert.Equal(t, len(value), valueSize)

	offset += VALUE_FIELD_SIZE
	extractedKey := string(encoded[offset : offset+keySize])
	assert.Equal(t, key, extractedKey)

	offset += keySize
	extractedValue := string(encoded[offset : offset+valueSize])
	assert.Equal(t, value, extractedValue)
}
