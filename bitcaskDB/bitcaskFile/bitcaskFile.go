package bitcaskFile

import "errors"

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
	return errors.New("bitcaskFile.Append not implemented")
}

func (bf *BitcaskFile) ReadAt(position int64, size int) ([]byte, error) {
	// Placeholder for read logic
	return nil, errors.New("bitcaskFile.ReadAt not implemented")
}

func (bf *BitcaskFile) CanAppend(dataSize int) bool {
	return true
}
