package reactive

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/reactivex/rxgo/v2"
)

func Test_RxGo001(t *testing.T) {
	observable := rxgo.Just("Hello, World!")()
	ch := observable.Observe()
	item := <-ch
	if item.Error() {
		fmt.Printf("error: %v \n", item.E)
	}
	fmt.Println("value is:", item.V)
}

func Test_RxGo002(t *testing.T) {
	ch := make(chan rxgo.Item)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- rxgo.Of(i)
		}
		close(ch)
	}()
	observable := rxgo.FromChannel(ch)

	// First Observer
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}

	// Second Observer
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}

func Test_RxGo003(t *testing.T) {
	observable := rxgo.Defer([]rxgo.Producer{func(_ context.Context, ch chan<- rxgo.Item) {
		for i := 0; i < 3; i++ {
			ch <- rxgo.Of(i)
		}
	}})

	// First Observer
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}

	// Second Observer
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}

func Test_RxGo004(t *testing.T) {
	ch := make(chan rxgo.Item)
	go func() {
		ch <- rxgo.Of(1)
		ch <- rxgo.Of(2)
		ch <- rxgo.Of(3)
		close(ch)
	}()
	// Create a Connectable Observable
	observable := rxgo.FromChannel(ch, rxgo.WithPublishStrategy())

	// Create the first Observer
	observable.DoOnNext(func(i interface{}) {
		fmt.Printf("First observer: %d\n", i)
	})

	// Create the second Observer
	observable.DoOnNext(func(i interface{}) {
		fmt.Printf("Second observer: %d\n", i)
	})

	disposedCtx, cancel := observable.Connect(context.Background())
	go func() {
		// Do something
		time.Sleep(time.Second)
		// Then cancel the subscription
		cancel()
	}()
	// Wait for the subscription to be disposed
	<-disposedCtx.Done()
}

func Test_RxGo005(t *testing.T) {
	observable := rxgo.Range(0, 3)

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}

func Test_RxGo006(t *testing.T) {
	observable := rxgo.Just(1, 2, 3)().
		Repeat(3, rxgo.WithDuration(1*time.Second))

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}

/*
  Hot Observable 例子
*/
func Test_RxGo007(t *testing.T) {
	ch := make(chan rxgo.Item)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- rxgo.Of(i)
		}
		close(ch)
	}()

	observable := rxgo.FromChannel(ch)

	for item := range observable.Observe() {
		fmt.Println("observe 1", item.V)
	}

	for item := range observable.Observe() {
		fmt.Println("observe 2", item.V)
	}
}

/*
  Cold Observable 例子
 */
func Test_RxGo008(t *testing.T) {
	observable := rxgo.Defer([]rxgo.Producer{func(_ context.Context, ch chan<- rxgo.Item) {
		for i := 0; i < 3; i++ {
			ch <- rxgo.Of(i)
		}
	}})

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}

func Test_RxGo009(t *testing.T) {
	observable := rxgo.Just(1, 2, 2, 3, 3, 4, 4)().
		Distinct(func(_ context.Context, i interface{}) (interface{}, error) {
			return i, nil
		})

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}

func Test_RxGo010(t *testing.T) {
	observable := rxgo.Just(1, 2, 3, 4, 5)().Skip(2)
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}

func Test_RxGo011(t *testing.T) {

}
