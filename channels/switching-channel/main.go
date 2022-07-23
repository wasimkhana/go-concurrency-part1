package main

import (
	"fmt"
)

func main() {

	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)

	select {
	case receivedMsg := <-msgCh:
		fmt.Println(receivedMsg)
	case receivedError := <-errCh:
		fmt.Println(receivedError)
	default:
		fmt.Println("No messages received")
	}

}

type Message struct {
	To      []string
	From    string
	Content string
}

type FailedMessage struct {
	ErrorMessage    string
	OriginalMessage Message
}

/*
Step-1 Demo setup
package main

import (
	"fmt"
)

func main() {

	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)

	msg := Message{
		To: []string{"simsim@example.com"},
		From: "simsim.org",
		Content: "Keep distance, Stay safe",
	}

	failedMessage := FailedMessage{
		ErrorMessage: "Message intercepted by black rider",
		OriginalMessage: Message{},
	}

	msgCh <- msg
	errCh <- failedMessage

	fmt.Println(<-msgCh)
	fmt.Println(<-errCh)

}

type Message struct {
	To []string
	From string
	Content string
}

type FailedMessage struct {
	ErrorMessage string
	OriginalMessage Message
}
*/

/*
Step-2 Add select, use first
package main

import (
	"fmt"
)

func main() {

	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)

	msg := Message{
		To: []string{"simsim@example.com"},
		From: "simsim.org",
		Content: "Keep distance, Stay safe",
	}

	msgCh <- msg

	select {
		case receivedMsg := <- msgCh:
			fmt.Println(receivedMsg)
		case receivedError := <- errCh:
			fmt.Println(receivedError)
	}

}

type Message struct {
	To []string
	From string
	Content string
}

type FailedMessage struct {
	ErrorMessage string
	OriginalMessage Message
}
*/

/*
Step-3 Use second select case
package main

import (
	"fmt"
)

func main() {

	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)

	msg := Message{
		To: []string{"simsim@example.com"},
		From: "simsim.org",
		Content: "Keep distance, Stay safe",
	}

	failedMessage := FailedMessage{
		ErrorMessage: "Message intercepted by black rider",
		OriginalMessage: Message{},
	}

	errCh <- failedMessage

	select {
		case receivedMsg := <- msgCh:
			fmt.Println(receivedMsg)
		case receivedError := <- errCh:
			fmt.Println(receivedError)
	}

}

type Message struct {
	To []string
	From string
	Content string
}

type FailedMessage struct {
	ErrorMessage string
	OriginalMessage Message
}
*/

/*
Step-4 Add select, send both messages
package main

import (
	"fmt"
)

func main() {

	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)

	msg := Message{
		To: []string{"simsim@example.com"},
		From: "simsim.org",
		Content: "Keep distance, Stay safe",
	}

	failedMessage := FailedMessage{
		ErrorMessage: "Message intercepted by black rider",
		OriginalMessage: Message{},
	}

	msgCh <- msg
	errCh <- failedMessage

	select {
		case receivedMsg := <- msgCh:
			fmt.Println(receivedMsg)
		case receivedError := <- errCh:
			fmt.Println(receivedError)
	}

}

type Message struct {
	To []string
	From string
	Content string
}

type FailedMessage struct {
	ErrorMessage string
	OriginalMessage Message
}
*/

/*
Step-5 Send no messages
package main

import (
	"fmt"
)

func main() {

	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)

	select {
		case receivedMsg := <- msgCh:
			fmt.Println(receivedMsg)
		case receivedError := <- errCh:
			fmt.Println(receivedError)
	}

}

type Message struct {
	To []string
	From string
	Content string
}

type FailedMessage struct {
	ErrorMessage string
	OriginalMessage Message
}
*/

/*
Step-6 Non-blocking select
package main

import (
	"fmt"
)

func main() {

	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)

	select {
		case receivedMsg := <- msgCh:
			fmt.Println(receivedMsg)
		case receivedError := <- errCh:
			fmt.Println(receivedError)
		default:
			fmt.Println("No messages received")
	}

}

type Message struct {
	To []string
	From string
	Content string
}

type FailedMessage struct {
	ErrorMessage string
	OriginalMessage Message
}
*/
