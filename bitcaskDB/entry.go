package bitcaskdb

type Entry struct {
	crc32     uint32
	keySize   int
	valueSize int
	key       string
	value     string
}

func NewEntry(key, value string) *Entry {
	return &Entry{
		crc32:     0, // Placeholder for CRC32 calculation
		keySize:   len(key),
		valueSize: len(value),
		key:       key,
		value:     value,
	}
}

func (e *Entry) Size() int {
	return 4 + 4 + 4 + e.keySize + e.valueSize // crc32 + keySize + valueSize + key + value
}
