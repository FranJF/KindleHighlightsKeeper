package htmltojson
import (
	"golang.org/x/net/html"
	"strings"
    "os"
)


var classNames = map[string]bool{
	"bookTitle":     true,
	"sectionHeading": true,
	"noteHeading":   true,
	"noteText":      true,
}

func ParseHTML(htmlFilePath string) (map[string][]string, string, error) {
	file, err := os.Open(htmlFilePath)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	doc, err := html.Parse(file)
	if err != nil {
		return nil, "", err
	}

	result := make(map[string][]string)
	var bookTitle string
	currentSection := ""
	var findClass func(*html.Node)
	findClass = func(n *html.Node) {
		switch {
		case n.Type == html.ElementNode && hasClass(n, classNames):
			className := getClass(n)
			if className == "bookTitle" {
				bookTitle = strings.TrimSpace(n.FirstChild.Data)
			} else if className == "sectionHeading" {
				currentSection = strings.TrimSpace(n.FirstChild.Data)
			} else if className == "noteText" && currentSection != "" {
				if n.FirstChild != nil {
					noteText := strings.TrimSpace(n.FirstChild.Data)
					result[currentSection] = append(result[currentSection], noteText)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findClass(c)
		}
	}
	findClass(doc)
	return result, bookTitle, nil
}

func hasClass(n *html.Node, classNames map[string]bool) bool {
	for _, attr := range n.Attr {
		if attr.Key == "class" {
			if classNames[attr.Val] {
				return true
			}
		}
	}
	return false
}

func getClass(n *html.Node) string {
	for _, attr := range n.Attr {
		if attr.Key == "class" {
			return attr.Val
		}
	}
	return ""
}
