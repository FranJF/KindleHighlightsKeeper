package htmltojson
import (
	"golang.org/x/net/html"
	"strings"
)


var classNames = map[string]bool{
	"bookTitle":     true,
	"sectionHeading": true,
	"noteHeading":   true,
	"noteText":      true,
}

func ParseHTML(htmlFilePath string) (string, map[string][]string, []string, error) {
	doc, err := html.Parse(strings.NewReader(htmlFilePath))
	if err != nil {
		return "", nil, nil, err
	}

	var bookTitle string
	currentSection := ""
    sections := []string{}
    data := make(map[string][]string)
	var findClass func(*html.Node)
    is_nota := false
	findClass = func(n *html.Node) {
		switch {
		case n.Type == html.ElementNode && hasClass(n, classNames):
			className := getClass(n)
			if className == "bookTitle" {
				bookTitle = strings.TrimSpace(n.FirstChild.Data)
			} else if className == "sectionHeading" {
				currentSection = strings.TrimSpace(n.FirstChild.Data)
                if currentSection != "" {
                    sections = append(sections, currentSection)
                }
            } else if className == "noteHeading"{
                heading := strings.TrimSpace(n.FirstChild.Data)
                if strings.Contains(heading, "Nota") {
                    is_nota = true
                }
            }else if className == "noteText" && currentSection != "" {
				if n.FirstChild != nil {
					noteText := strings.TrimSpace(n.FirstChild.Data)
                    if is_nota {
                        noteText = "Nota: " + noteText
                    }
                    is_nota = false
					data[currentSection] = append(data[currentSection], noteText)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findClass(c)
		}
	}
	findClass(doc)
	return bookTitle, data, sections, nil
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
