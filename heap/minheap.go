package heap

type MinHeap struct {
	data []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		data: []int{},
	}
}

func (m *MinHeap) swap(idx1, idx2 int) {
	m.data[idx1], m.data[idx2] = m.data[idx2], m.data[idx1]
}

func (m *MinHeap) shiftUp(idx int) {
	for idx > 0 {
		parent := (idx - 1) / 2

		if m.data[parent] > m.data[idx] {
			m.swap(parent, idx)
			idx = parent
		} else {
			break
		}
	}
}

func (m *MinHeap) Insert(val int) {
	m.data = append(m.data, val)
	m.shiftUp(len(m.data) - 1)
}
