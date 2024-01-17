package strct

import "errors"

type MeanHeap struct {
	length int
	data   []int
}

func (m *MeanHeap) Insert(a int) {
	m.data = append(m.data, a)
	m.heapifyUp(m.length)
	m.length++
}

func (m *MeanHeap) Delete() (int, error) {
	if m.length == 0 {
		return -1, errors.New("there is now data in heap")
	}

	value := m.data[0]
	m.length--
	m.data[0] = m.data[m.length]
	m.data = m.data[:m.length]
	if m.length != 0 {
		m.heapifyDown(0)
	}
	return value, nil
}

func (m *MeanHeap) heapifyUp(idx int) {
	if idx == 0 {
		return
	}

	parentIdx := m.parent(idx)
	if m.data[parentIdx] > m.data[idx] {
		m.data[parentIdx], m.data[idx] = m.data[idx], m.data[parentIdx]
		m.heapifyUp(parentIdx)
	}
}

func (m *MeanHeap) heapifyDown(idx int) {
	if idx >= m.length {
		return
	}

	leftIdx := m.leftChild(idx)
	if leftIdx >= m.length {
		return
	}

	rightIdx := m.rightChild(idx)
	if rightIdx >= m.length {
		if m.data[leftIdx] < m.data[idx] {
			m.data[leftIdx], m.data[idx] = m.data[idx], m.data[leftIdx]
		}
		return
	}

	if m.data[leftIdx] <= m.data[rightIdx] && m.data[leftIdx] < m.data[idx] {
		m.data[leftIdx], m.data[idx] = m.data[idx], m.data[leftIdx]
		m.heapifyDown(leftIdx)
	} else if m.data[rightIdx] < m.data[leftIdx] && m.data[rightIdx] < m.data[idx] {
		m.data[rightIdx], m.data[idx] = m.data[idx], m.data[rightIdx]
		m.heapifyDown(rightIdx)
	}
}

func (m MeanHeap) parent(idx int) int {
	return (idx - 1) / 2
}

func (m MeanHeap) leftChild(idx int) int {
	return idx*2 + 1
}

func (m MeanHeap) rightChild(idx int) int {
	return idx*2 + 2
}
