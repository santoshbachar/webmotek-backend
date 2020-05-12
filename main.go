package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/santoshbachar/webmotek-backend/engine"
)

func main() {
	fmt.Println("Hello world!")

	http.HandleFunc("/", HandleRequest)

	http.ListenAndServe(":8080", nil)
}

// HandleRequest handles incoming request
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "received the request /")

	if r.Method != http.MethodGet {
		fmt.Fprint(w, "Which idiot sent a POST request?")
		return
	}

	u := r.URL.Query()["url"]
	action := r.URL.Query()["action"]
	tags := r.URL.Query()["tags"]

	// fmt.Println(u, len(u))
	// fmt.Println(action, len(action))
	// fmt.Println(tags, len(tags))

	newURL, err := getParam(u)
	if err != nil {
		fmt.Fprint(w, "Sanitize error", err)
		return
	}
	fmt.Println(newURL)

	base, err := r.URL.Parse(newURL)
	if err != nil {
		fmt.Println("invalid url")
	}
	if base.Scheme == "" {
		fmt.Println("scheme is null. adding a http")
		newURL = "http://" + newURL
	}
	fmt.Println("base", base)

	// fetchDone := make(chan http.Response, r)
	// fmt.Println("base.String()", base.String())
	res := <-engine.FetchWebPage(newURL)
	defer res.Body.Close()

	strAction, err := getParam(action)
	if err != nil {
		fmt.Println("action param error. setting up the default value")
		strAction = ""
	}

	strTags, err := getParam(tags)
	if err != nil {
		fmt.Println("tags param error. setting up the default value")
		strTags = ""
	}

	/*parseDone := make(chan *goquery.Document, 1)
	go engine.Parse(parseDone, &res, strAction, strTags)
	*/

	doc := engine.Parse(&res, strAction, strTags)
	html, err := doc.Html()
	if err != nil {
		fmt.Fprint(w, "goquery to html error")
	}

	fmt.Println(doc.Html())

	fmt.Fprint(w, html)
}

// getParam returns the validated param
func getParam(url []string) (string, error) {
	if len(url) <= 0 {
		return "", errors.New("url<0")
	}
	str := url[0]

	return str, nil
}
