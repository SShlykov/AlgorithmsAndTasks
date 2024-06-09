package main

func main() {
	var counter int
	// Нет синхронизации (параллельно будет меняться одно и то же значение)
	for i := 0; i < 1000; i++ {
		go func() {
			counter++
		}()
	}
	// Нет ожидания запуска горутин, main рутина закончится раньше исполнения
}

// решается через sync.Mutex и sync.WaitGroup
