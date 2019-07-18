/**

We believe that coupling exceptions to a control structure, as in the try-catch-finally idiom, results in convoluted code.
It also tends to encourage programmers to label too many ordinary errors, such as failing to open a file, as exceptional.
https://golang.org/doc/faq#exceptions

The error type is an interface type. An error variable represents any value that can describe itself as a string.
Here is the interface's declaration:

type error interface {
    Error() string
}
https://blog.golang.org/error-handling-and-go

 */

package errorHandling

import (
	"fmt"
	"log"
)

type norgateMathError struct {
	lat  string
	long string
	err  error
}

func (n norgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occured: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("norgate math redux: square root of negative number: %v", f)
		return 0, norgateMathError{"50.2289 N", "99.4656 W", nme}
	}
	return 42, nil
}

// see use of structs with error type in standard library:
//
// http://golang.org/pkg/net/#OpError
// http://golang.org/src/pkg/net/dial.go
// http://golang.org/src/pkg/net/net.go
//
// http://golang.org/src/pkg/encoding/json/decode.go
