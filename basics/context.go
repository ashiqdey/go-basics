package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	timeoutContext, cancel := context.WithTimeout(context.Background(), time.Millisecond*2000)
	defer cancel()

	// create http request
	req, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, "http://placehold.it/2000x2000", nil)
	if err != nil {
		panic(err)
	}

	// perform http request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	// get data from HTTP response
	imageData, err := ioutil.ReadAll(res.Body)
	fmt.Printf("Image size %d\n", len(imageData))

}
