package database

type Index struct {
	key string
	blockSize int
	index map[string]int
}

func CreateIndex(key string, blockSize int) *Index {
	return &Index{ key: key, blockSize: blockSize }
}

func (i *Index) Search(key string) (int, error) {
	return 0, nil
}




