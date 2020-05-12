package engine

import (
	"log"
	"net/http"
)

// FetchWebPage downloads webpage
func FetchWebPage(url string) <-chan http.Response {

	resChan := make(chan http.Response, 1)

	go func() {

		res, err := http.Get(url)

		if err != nil {
			log.Fatal(err)
		}

		if res.StatusCode != 200 {
			log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
		}

		resChan <- *res

	}()

	return resChan

}
