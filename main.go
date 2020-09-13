package main

func work() {
	workers := 5
	ch := make(chan *Task, workers)
	defer close(ch)
	for i := 0; i < workers; i++ {
		go func() {
			for task := range ch {
				task.DoSomething()
			}
		}()
	}

	for i := 0; i < 20; i++ {
		ch <- &Task{}
	}
}
