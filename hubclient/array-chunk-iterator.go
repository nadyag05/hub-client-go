package hubclient

import "github.com/pkg/errors"

type ArrayChunkIterator struct {
	chunks   []string
	position int
}

func NewArrayChunkIterator(chunks []string) ArrayChunkIterator {
	i := new(ArrayChunkIterator)
	i.chunks = chunks
	i.position = -1
	return *i
}

func (i *ArrayChunkIterator) hasNext() bool {
	if len(i.chunks) > (i.position + 1) {
		return true
	} else {
		return false
	}
}

// next returns true and the next chunk if there is a next chunk. Returns false otherwise.
func (i *ArrayChunkIterator) next() (string, error) {
	i.position++
	if i.position < len(i.chunks) {
		return i.chunks[i.position], nil
	}
	return "", errors.New("A new chunk was requested but no chunks exist.")
}
