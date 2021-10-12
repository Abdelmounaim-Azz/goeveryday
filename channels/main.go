package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	var uri string
	var timeout time.Duration
	flag.StringVar(&uri, "uri", "", "uri to fetch")
	flag.DurationVar(&timeout, "timeout", time.Second*5, "timeout for HTTP client")
	flag.Parse()
	titleCh := make(chan string)
	errorCh := make(chan error)
	go titleOf(uri, timeout, titleCh, errorCh)
	select {
	case title := <-titleCh:
		fmt.Printf("Title: %s\n", title)
	case err := <-errorCh:
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
}
func titleOf(uri string, timeout time.Duration, titleCh chan string, errCh chan error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		errCh <- err
		return
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		errCh <- err
		return
	}
	if res.Body != nil {
		defer res.Body.Close()
		res, _ := ioutil.ReadAll(res.Body)
		for _, l := range strings.Split(string(res), "\n") {
			if strings.Contains(l, "<title>") {
				titleCh <- l
				return
			}
		}
	}
	errCh <- fmt.Errorf("no title found")
}
