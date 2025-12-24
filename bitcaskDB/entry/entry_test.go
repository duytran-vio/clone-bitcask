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
func TestNewEntry(t *testing.T) {
	t.Run("creates entry with valid key and value", func(t *testing.T) {
		key := "testKey"
		value := "testValue"
		entry := NewEntry(key, value)

		assert.NotNil(t, entry)
		assert.Equal(t, key, entry.key)
		assert.Equal(t, value, entry.value)
		assert.Equal(t, len(key), entry.keySize)
		assert.Equal(t, len(value), entry.valueSize)
		assert.Greater(t, entry.timestamp, int64(0))
	})

	t.Run("panics when key is empty", func(t *testing.T) {
		assert.Panics(t, func() {
			NewEntry("", "value")
		})
	})

	t.Run("panics when value is empty", func(t *testing.T) {
		assert.Panics(t, func() {
			NewEntry("key", "")
		})
	})

	t.Run("panics when both key and value are empty", func(t *testing.T) {
		assert.Panics(t, func() {
			NewEntry("", "")
		})
	})

	t.Run("panics when key exceeds maximum size", func(t *testing.T) {
		largeKey := string(make([]byte, KEY_MAX_CHARS+1))
		assert.Panics(t, func() {
			NewEntry(largeKey, "value")
		})
	})

	t.Run("panics when value exceeds maximum size", func(t *testing.T) {
		largeValue := string(make([]byte, VALUE_MAX_CHARS+1))
		assert.Panics(t, func() {
			NewEntry("key", largeValue)
		})
	})

	t.Run("accepts key at maximum size", func(t *testing.T) {
		maxKey := string(make([]byte, KEY_MAX_CHARS))
		entry := NewEntry(maxKey, "value")
		assert.NotNil(t, entry)
		assert.Equal(t, KEY_MAX_CHARS, entry.keySize)
	})

	t.Run("accepts value at maximum size", func(t *testing.T) {
		maxValue := string(make([]byte, VALUE_MAX_CHARS))
		entry := NewEntry("key", maxValue)
		assert.NotNil(t, entry)
		assert.Equal(t, VALUE_MAX_CHARS, entry.valueSize)
	})
}

