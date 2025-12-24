package bitcaskFile

type BitcaskFile struct {
	fileID   int
	filePath string
	size     int64
}

func NewBitcaskFile(fileID int, filePath string) *BitcaskFile {
	return &BitcaskFile{
		fileID:   fileID,
		filePath: filePath,
		size:     0,
	}
}

func (bf *BitcaskFile) Size() int64 {
	return bf.size
}

func (bf *BitcaskFile) FileID() int {
	return bf.fileID
}

func (bf *BitcaskFile) Append(data []byte) error {
	// Placeholder for write logic
	return nil
}

func (bf *BitcaskFile) Put(data []byte) error {
	// Placeholder for write logic
	return nil
}

func (bf *BitcaskFile) ReadAt(position int64, size int) ([]byte, error) {
	// Placeholder for read logic
	return nil, nil
}

func (bf *BitcaskFile) CanAppend(dataSize int) bool {
	return true
}
