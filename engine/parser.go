package engine

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func defaultParser(doc *goquery.Document) {
	RemoveImage(doc)
	RemoveSvg(doc)
	RemoveVideo(doc)

	RemoveScript(doc)
	RemoveLink(doc)

	RemoveForm(doc)
	RemoveInput(doc)
	RemoveButton(doc)
	RemoveTextArea(doc)
}

// SearchParser ...
func SearchParser(res *http.Response, action string, tags string) []SearchModel {

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("about to do defaultParser")

	// defaultParser(doc)
	/*doc.Find("head").Each(func(index int, ele *goquery.Selection) {
		// head is only left with title and style, so
		// it can be remove in future
		// ele.Remove()
		style := ele.Find("style")
		style.Remove()
	})*/

	fmt.Println("default parser done")

	var search []SearchModel

	// This one works with the desktop version of google
	// changing the agent will do
	// but is bulky than we got in mobile version
	// so we are not going to use this version
	// and instead use the mobile version
	/*doc.Find(".search").Each(func(i int, s *goquery.Selection) {
		fmt.Println("enter .search")
		r := s.Find(".r")
		a := r.Find("a")
		href, exists := a.Attr("href")
		if exists == false {
			fmt.Println("no href")
		}
		h3 := r.Find("h3").Text()

		// description
		d := s.Find(".s").Find(".st").Text()

		fmt.Println("title", h3)
		fmt.Println("a", href)
		fmt.Println("description", d)

		search = append(search, SearchModel{title: h3, URL: href, description: d})

		fmt.Println("search after append", search)
	})*/

	// equivalent of .search in desktop mode
	fmt.Println("going to do doc.Find")
	doc.Find("div.ZINbbc.xpd.O9g5cc.uUPGi").Each(func(i int, s *goquery.Selection) {
		// fmt.Println("enter .ZINbbc xpd O9g5cc uUPGi")
		// sm := s.Find("div.ZINbbc.xpd.O9g5cc.uUPGi")
		r := s.Find("div.kCrYT") // equivalent of .r in desktop mode
		a := r.Find("a")
		href, exists := a.Attr("href")
		if exists == false {
			fmt.Println("no href")
			return
		}

		// in mobile version just getting the href is not done yet.
		// we have to parse it to get the url inside of it.

		hrefParse, err := url.Parse(href)
		if err != nil {
			fmt.Println("href parse error", err)
			return
		}

		hrefValues, err := url.ParseQuery(hrefParse.RawQuery)
		if err != nil {
			fmt.Println("hrefValues error", err)
			return
		}

		// q is the new href
		q := hrefValues.Get("q")

		// here in mobile version title is there in ".a" instead of ".r"
		h3 := a.Find("div.BNeawe.vvjwJb.AP7Wnd").Text()

		// description
		// .kCrYT instead of .s
		// BNeawe s3v9rd AP7Wnd instead of .st
		d := s.Find("div.kCrYT").Find("div.BNeawe.s3v9rd.AP7Wnd").Text()

		fmt.Println("title", h3)
		//fmt.Println("a", q)
		fmt.Println("description", d)

		search = append(search, SearchModel{Title: h3, URL: q, Description: d})

		// fmt.Println("search after append", search)
	})
	fmt.Println("doc.Find() done")
	fmt.Println("len of search", len(search))
	return search
}

func allowParser(doc **goquery.Document, tags *[]string) {

	for _, tag := range *tags {
		switch tag {
		case "img":
			if !containsTag(tags, tag) {
				RemoveImage(*doc)
			}
			break
		}
	}
}

func blockParser(doc *goquery.Document, tags *[]string) {

	for _, tag := range *tags {
		switch tag {
		case "img":
			if containsTag(tags, tag) {
				RemoveImage(doc)
			}
			break
		case "svg":
			if containsTag(tags, tag) {
				RemoveSvg(doc)
			}
			break
		case "video":
			if containsTag(tags, tag) {
				RemoveVideo(doc)
			}
			break
		case "script":
			if containsTag(tags, tag) {
				RemoveScript(doc)
			}
			break
		case "link":
			if containsTag(tags, tag) {
				RemoveLink(doc)
			}
			break
		case "form":
			if containsTag(tags, tag) {
				RemoveForm(doc)
			}
			break
		case "input":
			if containsTag(tags, tag) {
				RemoveInput(doc)
			}
			break
		case "button":
			if containsTag(tags, tag) {
				RemoveTextArea(doc)
			}
			break
		case "textarea":
			if containsTag(tags, tag) {
				RemoveTextArea(doc)
			}
			break
		}
	}
}

// Parse - entry point of
func Parse(res *http.Response, action string, tags string) *goquery.Document {

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	slcTags := strings.Split(tags, ",")

	switch action {
	case "allow":
		allowParser(&doc, &slcTags)
		break
	case "block":
		break
	default:
		defaultParser(doc)
	}

	fmt.Println("Parse completed")

	return doc

}

func containsTag(slc *[]string, ele string) bool {
	for _, v := range *slc {
		if v == ele {
			return true
		}
	}
	return false
}

// RemoveImage removes img tag element
func RemoveImage(doc *goquery.Document) {
	doc.Find("img").Each(func(index int, ele *goquery.Selection) {
		ele.Remove()
	})
}

// RemoveSvg removes svg tag element
func RemoveSvg(doc *goquery.Document) {
	doc.Find("svg").Each(func(index int, ele *goquery.Selection) {
		ele.Remove()
	})
}

// RemoveVideo removes video tag element
func RemoveVideo(doc *goquery.Document) {
	doc.Find("video").Each(func(index int, ele *goquery.Selection) {
		ele.Remove()
	})
}

// RemoveScript removes script tag element
func RemoveScript(doc *goquery.Document) {
	doc.Find("script").Each(func(index int, ele *goquery.Selection) {
		ele.Remove()
	})
}

// RemoveLink removes link tag element
func RemoveLink(doc *goquery.Document) {
	doc.Find("link").Each(func(index int, ele *goquery.Selection) {
		ele.Remove()
	})
}

// RemoveForm removes form tag element
func RemoveForm(doc *goquery.Document) {
	doc.Find("form").Each(func(index int, ele *goquery.Selection) {
		ele.Remove()
	})
}

// RemoveInput removes input tag element
func RemoveInput(doc *goquery.Document) {
	doc.Find("input").Each(func(index int, ele *goquery.Selection) {
		ele.Remove()
	})
}

// RemoveButton removes input button element
func RemoveButton(doc *goquery.Document) {
	doc.Find("button").Each(func(index int, ele *goquery.Selection) {
		ele.Remove()
	})
}

// RemoveTextArea removes textarea tag element
func RemoveTextArea(doc *goquery.Document) {
	doc.Find("textarea").Each(func(index int, ele *goquery.Selection) {
		ele.Remove()
	})
}
