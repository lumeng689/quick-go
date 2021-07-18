package main
import (
	"time"
	"fmt"
	"runtime"
	"sync"
)

func main() {
	for i := 0; i < 10; i ++ {
		go func() {
			fmt.Printf("%d is running.\n", i);
			time.Sleep(1 * time.Second)
			fmt.Printf("%d is stopping.\n", i);
		}();
		time.Sleep(time.Second);
	}

	names := []string{"Eric", "Harry", "Robert", "Jim", "Mark"}
	for _, name := range names {
		go func() {
			fmt.Printf("Hello,%s.\n", name)
		}()
		runtime.Gosched()
	}

	for _, name := range names {
		go func(who string) {
			fmt.Printf("Hello,%s.\n", who)
		}(name)
	}

	runtime.Gosched()

	// select test
	var ch1 = make(chan int, 10)
	var ch2 = make(chan int, 10)

	select {
	case e1 := <-ch1:
		fmt.Printf("1th case is selected . e1=%v.\n", e1)
	case e2 := <-ch2:
		fmt.Printf("2th case is selected. e2=%v,\n", e2)
	default:
		fmt.Println("default!")
	}

	repeatedLock()
}

func repeatedLock() {
	var mutex sync.Mutex
	fmt.Println("Lock the lock.(G0)")
	mutex.Lock()
	fmt.Println("The lock is locked.(G0)")
	for i:=1;i<=3;i++ {
		go func(i int) {
			fmt.Printf("Lock the lock.(G%d)\n",i)
			mutex.Lock()
			fmt.Printf("The lock is locked(G%d)\n",i)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("Unlock the lock.(G0)")
	mutex.Unlock()
	fmt.Println("The lock is unlocked.(G0)")
	time.Sleep(time.Second)
}
