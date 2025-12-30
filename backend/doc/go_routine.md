```go
go func()
<-channel // รอ
```

```go
func sayHello(done chan bool) {
	fmt.Println("Hello")
	done <- true // ส่งสัญญาณว่าทำเสร็จแล้ว
}

func main() {
	done := make(chan bool)

	go sayHello(done)

	<-done // รอจนกว่าจะมีค่าเข้ามา
	fmt.Println("Done")
}
```

sync.Waitgroup

```go
var wg sync.WaitGroup

func worker(id int) {
	defer wg.Done()
	fmt.Println("Worker", id)
}

func main() {
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i)
	}

	wg.Wait()
	fmt.Println("All done")
}

```

```go
    jobs := make(chan int)
	results := make(chan *m.FixturePredictionBundle)
	errCh := make(chan error, 1)
//ดึง worker ตาม cpu core
workers := runtime.NumCPU()
```
