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

func (m *MinHeap) shiftDown(i int) {
	smallest := i
	length := len(m.data)
	left, right := (i*2)+1, (i*2)+2

	for {
		if left < length && m.data[smallest] > m.data[left] {
			smallest = left
		}
		if right < length && m.data[smallest] > m.data[right] {
			smallest = right
		}

		if smallest == i {
			break
		}
		m.swap(smallest, i)
		i = smallest
	}
}

func (m *MinHeap) Insert(val int) {
	m.data = append(m.data, val)
	m.shiftUp(len(m.data) - 1)
}

// extract the root element (min value) and readjust the nodes
// to maintain min heap
func (m *MinHeap) Extract() int {
	root := m.data[0]

	lastIdx := len(m.data) - 1
	m.data[0] = m.data[lastIdx]
	m.data = m.data[:lastIdx]

	m.shiftDown(0)

	return root
}
