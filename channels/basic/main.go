package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 1)

	ch <- "Hello"

	fmt.Println(<-ch)

	ch <- "Go"

	fmt.Println(<-ch)
}

/*
Step-1 Basic channel with no buffer
package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	ch <- "Hello"

	fmt.Println(<-ch)
}
*/

/*
Step-2 Add buffer to channel to make non-blocking on one message
package main

import (
	"fmt"
)

func main() {
	ch := make(chan string,1)

	ch <- "Hello"

	fmt.Println(<-ch)
}
*/

/*
Step-3 Send and receive second message in channel

package main

import (
	"fmt"
)

func main() {
	ch := make(chan string,1)

	ch <- "Hello"

	fmt.Println(<-ch)

	ch <- "Go"

	fmt.Println(<-ch)
}
*/
