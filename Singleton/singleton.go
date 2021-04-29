package singleton

import "sync"

var once sync.Once
var mu sync.Mutex

type Singleton interface {
	AddOne() int
	GetCount() int
}

type singleton struct {
	count int
	sync.RWMutex
}

func (s *singleton) AddOne() int {
	s.Lock()
	defer s.Unlock()

	s.count++
	return s.count
}

func (s *singleton) GetCount() int {
	s.RLock()
	defer s.RUnlock()

	return s.count
}

var instance *singleton

func GetInstance() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
