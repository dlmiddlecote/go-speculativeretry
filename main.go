package speculativeretry

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Starting")

	// response channel
	respCh := make(chan *http.Response)
	// err channel
	errCh := make(chan error)

	// do request
	go func(respCh chan *http.Response, errCh chan error) {
		time.Sleep(1 * time.Second)
		resp, err := http.Get("http://httpbin.org/ip")
		if err != nil {
			errCh <- err
		} else {
			respCh <- resp
		}
	}(respCh, errCh)

	select {
	case resp := <-respCh:
		fmt.Println(resp)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("Do request again")
	case err := <-errCh:
		fmt.Println("Err:", err)
	}

	fmt.Println("Done!")
}
