package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinHeap(t *testing.T) {
	t.Run("Insert", testInsert)
	t.Run("Extract", testExtract)
}

func getMinHeap() *MinHeap {
	h := NewMinHeap()
	h.Insert(20)
	h.Insert(2)
	h.Insert(34)
	h.Insert(9)
	h.Insert(1)
	h.Insert(8)
	h.Insert(5)

	return h
}

func testInsert(t *testing.T) {

	h := getMinHeap()
	l := len(h.data)
	for i := 0; i < l; i++ {
		left, right := (i*2)+1, (i*2)+2

		if left < l && h.data[left] < h.data[i] {
			t.Errorf("Heap property violated at index %d, %d > left %d", i, h.data[i], h.data[left])
		}
		if right < l && h.data[right] < h.data[i] {
			t.Errorf("Heap property violated at index %d, %d > right %d", i, h.data[i], h.data[right])
		}

	}
}

func testExtract(t *testing.T) {

	h := getMinHeap()

	assert.Equal(t, 1, h.Extract())
	assert.Equal(t, 2, h.Extract())
}
