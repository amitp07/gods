package heap

import (
	"testing"
)

func TestMinHeap(t *testing.T) {
	t.Run("Insert", testInsert)
}

func testInsert(t *testing.T) {
	h := NewMinHeap()
	h.Insert(20)
	h.Insert(2)
	h.Insert(34)
	h.Insert(9)
	h.Insert(1)
	h.Insert(8)
	h.Insert(5)

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
