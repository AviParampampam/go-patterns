package singleton

import (
	"fmt"
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T) {
	var counter1 Singleton = GetInstance()
	if counter1 == nil {
		t.Fatalf("Ожидался указатель на Singleton, а не nil")
	}

	currentCount := counter1.AddOne()
	if currentCount != 1 {
		t.Errorf("После первого вызова для подсчета счета должен быть 1, а он %d\n", currentCount)
	}

	var expectedCounter Singleton = counter1
	var counter2 Singleton = GetInstance()
	if counter2 != expectedCounter {
		t.Error("Ожидается такой же экземпляр в сounter2, но у него другой экземпляр")
	}

	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Errorf("После вызова AddOne с использованием второго счетчика текущий счет должен быть равен 2, но был %d\n", currentCount)
	}
}

func TestParallel(t *testing.T) {
	singleton := GetInstance()
	singleton2 := GetInstance()

	n := 5000

	var wg sync.WaitGroup

	fmt.Printf("До цикла текущий счетчик %d\n", singleton.GetCount())

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			singleton.AddOne()
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			singleton2.AddOne()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("Текущее количество %d\n", singleton.GetCount())

	currentCount1 := singleton.GetCount()
	currentCount2 := singleton2.GetCount()
	if currentCount1 != currentCount2 {
		t.Errorf("Счетчики не совпадают\nCurrentCount1=%d\nCurrentCount2=%d", currentCount1, currentCount2)
	}

	if currentCount1 != n*2 {
		t.Errorf("Количество не совпадает\nCurrentCount1=%d\nN*2=%d", currentCount1, n*2)
	}
}
