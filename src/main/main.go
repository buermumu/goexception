package main

import (
	"exception"
	"fmt"
	"time"
)

func handler(message interface{}, args ...interface{}) {
	// file.WriteString(fmt.Sprintf("msg:%s", message))
	fmt.Println(fmt.Sprintf("message:%s", message))
	fmt.Println("args:", args)
}

func task_1() {
	fmt.Println("task_1 runing")
	panic("task_1 exception")
	fmt.Println("task_1 done")
}

func task_2() {
	h := exception.New()
	defer h.Catch(handler, "c")
	h.Try(func() {
		fmt.Println("task_2 runing")
		panic("task_2 exception")
		fmt.Println("task_2 done")
	})
}

func main() {
	// test_1
	h := exception.New()
	defer h.Catch(handler, "a", "b")
	h.Try(func() {
		task_1()
	})

	// test_2
	go task_2()

	// test_3
	h = exception.New()
	defer h.Catch(handler)
	h.Try(func() {
		h.Throw("task_3 exception")
	})

	// test_4
	a := exception.New()
	defer a.Catch(handler)
	a.Try(func() {
		fmt.Println("aa")
		b := exception.New()
		defer b.Catch(handler)
		b.Try(func() {
			fmt.Println("bb")
			b.Throw("bb exception")
		})
		a.Throw("aa exception")
	})

	time.Sleep(time.Duration(5) * time.Second)
}
