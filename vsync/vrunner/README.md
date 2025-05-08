vrunner - run/stop goroutines/tasks securely, recursively.


```go
s1 := vrunner.New()

s1.Defer(func() {
    fmt.Println("s1 stopped 2")
})
s1.Defer(func() {
    fmt.Println("s1 stopped 1")
})

// loop run task until s1 closed.
s1.Loop(func() {
    fmt.Println("s1 run loop task")
    time.Sleep(time.Millisecond*3)
})

// run interval task until s1 closed.
s1.Interval(func() {
    t.Log("s1 run interval task")
}, time.Millisecond)

go func() {
    ticker := time.NewTicker(time.Millisecond*2)
    
    for {
        select {
        case <-s1.C:
            return
        case <-ticker.C:
        	fmt.Println("run ticker task until s1 stopped")
        }
    }
}()

s2 := s1.NewChild()
s2.Defer(func() {
    fmt.Println("s2 stopped")
})

s3 := s2.NewChild()
s3.Defer(func() {
    fmt.Println("s3 stopped")
})

time.Sleep(time.Millisecond * 10)

s1.Stop()

time.Sleep(time.Millisecond * 10)

// s1 run loop task
// s1 run interval task
// s1 run interval task
// run ticker task until s1 stopped
// s1 run interval task
// s1 run loop task
// s1 run interval task
// run ticker task until s1 stopped
// s1 run interval task
// s1 run interval task
// run ticker task until s1 stopped
// s1 run loop task
// s1 run interval task
// s1 run interval task
// run ticker task until s1 stopped
// s1 run interval task
// s1 run loop task
// s1 run interval task
// s1 stopped 1
// s3 stopped
// s2 stopped
// s1 stopped 2
```