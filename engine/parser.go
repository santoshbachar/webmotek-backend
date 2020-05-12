package engine

import (
	"fmt"
	"log"
	"net/http"
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
