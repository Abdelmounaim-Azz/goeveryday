package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	uris := []string{"https://github.com/Abdelmounaim-Azz",
		"https://github.com/docker",
		"https://github.com/kubernetes",
		"https://github.com/moby",
		"https://github.com/hashicorp",
	}
	var w int
	flag.IntVar(&w, "w", 1, "amount of workers")
	flag.Parse()
	titleMap, errs := fetch(uris, w)
	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Fprintf(os.Stderr, "error %s", err.Error())
		}
		os.Exit(1)
	}
	for k, v := range titleMap {
		fmt.Printf("%s has %s\n", k, v)
	}
}
func fetch(uris []string, workers int) (map[string]string, []error) {
	titles := make(map[string]string)
	var errs []error
	workQueue := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		worker := i
		go func(worker int, workQueue chan string) {
			log.Printf("[Worker: %d] Started work.", worker)
			for uri := range workQueue {
				log.Printf("[Worker: %d] Getting: %s", worker, uri)
				start := time.Now()
				title, err := titleOf(uri, time.Second*5)
				if err != nil {
					errs = append(errs, err)
				}
				log.Printf("[Worker: %d] Got: %s (%.2fs)",
					worker,
					uri,
					time.Since(start).Seconds())
				titles[uri] = title
			}
			log.Printf("[Worker: %d] Finished work.", worker)
			wg.Done()
		}(worker, workQueue)
	}
	go func() {
		for _, u := range uris {
			workQueue <- u
		}
		close(workQueue)
	}()
	wg.Wait()
	return titles, errs
}

func titleOf(uri string, timeout time.Duration) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return "", err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	if res.Body != nil {
		defer res.Body.Close()
		res, _ := ioutil.ReadAll(res.Body)
		for _, l := range strings.Split(string(res), "\n") {
			if strings.Contains(l, "<title>") {
				return l, nil
			}
		}
	}
	return "", fmt.Errorf("no title found")
}
