package engine

import (
	"log"
	"net/http"
)

// ScrapeGoogle fetches Google Search Result in desktop mode
func ScrapeGoogle(url string) <-chan http.Response {
	resChan := make(chan http.Response, 1)

	go func() {

		client := &http.Client{}

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatalln("ScrapeGoogle() NewRequest error", err)
		}

		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.141 Safari/537.36")

		res, err := client.Do(req)
		if err != nil {
			log.Fatalln("client.Do error", err)
		}

		// This is not required since we want *Response
		/*body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalln("ioutil.ReadAll error", err)
		}*/

		resChan <- *res

	}()

	return resChan

}
