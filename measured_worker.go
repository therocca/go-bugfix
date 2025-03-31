package main

import (
	"sync/atomic"
)

type MeasuredWorker struct {
	Worker
	value  int32
}

func (m *MeasuredWorker) Work() {
	m.Worker.Work()
	atomic.AddInt32(&m.value, 1)
}

func (m *MeasuredWorker) Value() int {
	return int(atomic.LoadInt32(&m.value))
}
