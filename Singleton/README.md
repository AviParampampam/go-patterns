# Паттерн «Одиночка»

**Одиночка** (англ. Singleton) — шаблон проектирования, гарантирующий, что у типа есть только один экземпляр, предоставляющий глобальную точку доступа к нему.

## Реализация (см. [код](./singleton.go))
```go
type Singleton interface {
  AddOne() int
}

type singleton struct {
  count int
}

var instance *singleton

func GetInstance() Singleton {
  if instance == nil {
    instance = &singleton{}
  }
  return instance
}

func (s *singleton) AddOne() int {
  s.count++
  return s.count
}
```
Проблема этой реализации становится сразу очевидна, функция `GetInstance` и метод `AddOne` подвержены ошибкам конкурентного доступа.

### Исправление ошибки

Сначала разберемся с `GetInstance`. Здесь можно попробовать исправить эту ошибку с помощью блокировок
```go
var mu sync.Mutex

func GetInstance() Singleton {
  mu.Lock()
  defer mu.Unlock

  if instance == nil {
    instance = &singleton{}
  }
  return instance
}
```

Но в стандартной библиотеке Golang есть `sync.Once`, который способен вызывать
функцию только один раз
```go
var once sync.Once

func GetInstance() Singleton {
  once.Do(func() {
    instance = &singleton{}
  })

  return instance
}
```

Теперь `AddOne`. Есть сейчас запустить текст `TestParallel` из файла [singleton_test.go](singleton_test.go), то ожидая получить 10000, мы получим гораздо меньше. Если запустить тест с `-race`, то мы увидем множество *DATA RACE* ошибок.

Чтобы исправить это, мы можем использовать бинарный семафор, он же *мьютекс*

```go
type singleton struct {
   count int
   sync.RWMutex
}

func (s *singleton) AddOne() {
   s.Lock()
   defer s.Unlock()
   s.count++
}

func (s *singleton) GetCount()int {
   s.RLock()
   defer s.RUnlock()
   return s.count
}
```

## Дополнительная информация
Все материалы взяты [отсюда](https://medium.com/german-gorelkin/go-singleton-f408a6c11a55)