package main

import (
	"fmt"
	"sync"
)

/*
Реализовать безопасную для конкуренции запись данных в структуру map.
Подсказка: необходимость использования синхронизации (например, sync.Mutex или встроенная concurrent-map).
Проверьте работу кода на гонки (util go run -race).
*/

type SyncMap struct {
	data map[int]int
	mu   sync.Mutex
	wg   *sync.WaitGroup
}

func NewSyncMap(wg *sync.WaitGroup) *SyncMap {
	return &SyncMap{data: make(map[int]int), wg: wg}
}

func (m *SyncMap) Add(key, val int) {
	defer m.wg.Done()
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = val
}

func (m *SyncMap) Print() {
	fmt.Println(m.data)
}

func main() {
	var wg sync.WaitGroup
	m := NewSyncMap(&wg)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go m.Add(i+1, i)
		wg.Add(1)
		go m.Add(i, i)
	}
	wg.Wait()
	m.Print()
}
