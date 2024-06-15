package channels

// channels used to send and transfer data between goroutines
// important for achieving concurrency in go routine
// same channel -> for 2 go routines to communicate -> bidirectional workflow
//
type userInput interface{}

// return will be an empty interface that might return channel of any compatibilty
func CreateChannelAnyDataype(input userInput) interface{} {

	ch := make(chan interface{})

	go func() {
		ch <- input
		close(ch)
	}()

	return ch
}
